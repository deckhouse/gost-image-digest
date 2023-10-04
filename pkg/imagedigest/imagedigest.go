package imagedigest

import (
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"github.com/rs/zerolog/log"
	"go.cypherpunks.ru/gogost/v5/gost34112012256"
)

const (
	gostDigestAnnotationKey = "deckhouse.io/gost-digest"
	registryTimeout         = 30 * time.Minute
)

type ImageMetadata struct {
	ImageName       string
	ImageDigest     string
	ImageGostDigest string
	LayersDigest    []string
}

func CalculateGostImageDigest(imageName string, opts ...crane.Option) ([]byte, error) {
	im, err := getImageMetadataFromRegistry(imageName, opts...)
	if err != nil {
		return nil, err
	}

	return calculateLayersGostDigest(im)
}

func AddGostImageDigest(imageName string, opts ...crane.Option) error {
	image, err := getImageFromRegistry(imageName, opts...)
	if err != nil {
		return err
	}
	im, err := imageToImageMetadata(imageName, image)
	if err != nil {
		return err
	}

	gostImageDigest, err := calculateLayersGostDigest(im)
	if err != nil {
		return err
	}
	log.Info().Msgf("GOST Image Digest: %s", hex.EncodeToString(gostImageDigest))
	return updateImageInRegistry(
		imageName,
		image,
		map[string]string{
			gostDigestAnnotationKey: hex.EncodeToString(gostImageDigest),
		},
		opts...,
	)
}

func ValidateGostImageDigest(imageName string, opts ...crane.Option) error {
	im, err := getImageMetadataFromRegistry(imageName, opts...)
	if err != nil {
		return err
	}

	if len(im.ImageGostDigest) == 0 {
		return fmt.Errorf("The image %s does not contain Gost Image Digest", imageName)
	}

	log.Info().Msgf("GOST Image Digest from image %s", im.ImageGostDigest)

	gostImageDigest, err := calculateLayersGostDigest(im)
	if err != nil {
		return err
	}
	log.Info().Msgf("Calculated GOST Image Digest %s", hex.EncodeToString(gostImageDigest))

	return compareImageGostHash(im, gostImageDigest)
}

func getImageMetadataFromRegistry(imageName string, opts ...crane.Option) (*ImageMetadata, error) {
	image, err := getImageFromRegistry(imageName, opts...)
	if err != nil {
		return nil, err
	}
	return imageToImageMetadata(imageName, image)
}

func getImageFromRegistry(imageName string, opts ...crane.Option) (v1.Image, error) {
	return crane.Pull(imageName, opts...)
}

func updateImageInRegistry(
	imageName string,
	image v1.Image,
	annotations map[string]string,
	opts ...crane.Option,
) error {
	image = mutate.Annotations(image, annotations).(v1.Image)
	return crane.Push(image, imageName, opts...)
}

func imageToImageMetadata(imageName string, image v1.Image) (*ImageMetadata, error) {
	im := &ImageMetadata{ImageName: imageName}

	imageDigest, err := image.Digest()
	if err != nil {
		return nil, err
	}
	im.ImageDigest = imageDigest.String()

	manifest, err := image.Manifest()
	if err != nil {
		return nil, err
	}

	imageGostDigestStr, ok := manifest.Annotations[gostDigestAnnotationKey]
	if !ok {
		log.Debug().Msg("the image does not contain gost digest")
	}
	im.ImageGostDigest = imageGostDigestStr

	layers, err := image.Layers()
	if err != nil {
		return nil, err
	}

	for _, layer := range layers {
		digest, err := layer.Digest()
		if err != nil {
			return nil, err
		}
		im.LayersDigest = append(im.LayersDigest, digest.String())
	}

	sort.Slice(
		im.LayersDigest,
		func(i, j int) bool {
			if strings.Compare(im.LayersDigest[i], im.LayersDigest[j]) == -1 {
				return true
			}
			return false
		},
	)

	log.Debug().Interface("im", im).Msg("ImageManifest")

	return im, nil
}

func calculateLayersGostDigest(im *ImageMetadata) ([]byte, error) {
	layersDigestBuilder := strings.Builder{}
	for _, digest := range im.LayersDigest {
		layersDigestBuilder.WriteString(digest)
	}

	data := layersDigestBuilder.String()

	if len(data) == 0 {
		return nil, fmt.Errorf("invalid layers hash data")
	}

	hasher := gost34112012256.New()
	_, err := hasher.Write([]byte(data))
	if err != nil {
		return nil, err
	}

	return hasher.Sum(nil), nil
}

func mutateConfig(base v1.Image) (v1.Image, error) {
	cfg, err := base.ConfigFile()
	if err != nil {
		return nil, err
	}

	if cfg.Config.Labels == nil {
		cfg.Config.Labels = map[string]string{}
	}

	base, err = mutate.ConfigFile(base, cfg)
	if err != nil {
		return nil, err
	}

	return base, nil
}

func compareImageGostHash(im *ImageMetadata, gostHash []byte) error {
	imageGostHashByte, err := hex.DecodeString(im.ImageGostDigest)
	if err != nil {
		return fmt.Errorf("invalid gost image digest: %w", err)
	}

	if subtle.ConstantTimeCompare(imageGostHashByte, gostHash) == 0 {
		return fmt.Errorf("invalid gost image digest comparation")
	}
	return nil
}

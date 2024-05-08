package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/labstack/echo/v4"
)

// TODO: used by containers

// models
type imageModel struct {
	// holds digests of image manifests that reference the image.
	ID string

	// RepoTags is a list of image names/tags in the local image cache that
	// reference this image.
	//
	// Multiple image tags can refer to the same image, and this list may be
	// empty if no tags reference the image, in which case the image is
	// "untagged", in which case it can still be referenced by its ID.
	RepoTags []string

	// RepoDigests is a list of content-addressable digests of locally available
	// image manifests that the image is referenced from. Multiple manifests can
	// refer to the same image.
	//
	// These digests are usually only available if the image was either pulled
	// from a registry, or if the image was pushed to a registry, which is when
	// the manifest is generated and its digest calculated.
	// RepoDigests []string

	// Parent is the ID of the parent image.
	//
	// Depending on how the image was created, this field may be empty and
	// is only set for images that were built/created locally. This field
	// is empty if the image was pulled from an image registry.
	Parent string

	// Comment is an optional message that can be set when committing or
	// importing the image.
	Comment string

	// Created is the date and time at which the image was created, formatted in
	// RFC 3339 nano-seconds (time.RFC3339Nano).
	//
	// This information is only available if present in the image,
	// and omitted otherwise.
	Created string

	// DockerVersion is the version of Docker that was used to build the image.
	//
	// Depending on how the image was created, this field may be empty.
	// DockerVersion string

	// Author is the name of the author that was specified when committing the
	// image, or as specified through MAINTAINER (deprecated) in the Dockerfile.
	// Author string
	// Config *container.Config

	// Architecture is the hardware CPU architecture that the image runs on.
	// Architecture string

	// Variant is the CPU architecture variant (presently ARM-only).
	// Variant string `json:",omitempty"`

	// OS is the Operating System the image is built to run on.
	// Os string

	// OsVersion is the version of the Operating System the image is built to
	// run on (especially for Windows).
	// OsVersion string `json:",omitempty"`

	// Size is the total size of the image including all layers it is composed of.
	Size int64

	// VirtualSize is the total size of the image including all layers it is
	// composed of.
	//
	// Deprecated: this field is omitted in API v1.44, but kept for backward compatibility. Use Size instead.
	// VirtualSize int64 `json:"VirtualSize,omitempty"`

	// GraphDriver holds information about the storage driver used to store the
	// container's and image's filesystem.
	// GraphDriver GraphDriverData

	// RootFS contains information about the image's RootFS, including the
	// layer IDs.
	// RootFS RootFS

	// Metadata of the image in the local cache.
	//
	// This information is local to the daemon, and not part of the image itself.
	// Metadata image.Metadata
}

type imageHistoryModel struct {

	// comment
	// Required: true
	// Comment string `json:"Comment"`

	// created
	// Required: true
	Created int64

	// created by
	// Required: true
	// CreatedBy string `json:"CreatedBy"`

	// Id
	// Required: true
	ID string

	// size
	// Required: true
	Size int64

	// tags
	// Required: true
	Tags []string
}

type imagePageModel struct {
	Image   imageModel
	History []imageHistoryModel
}

// handler
func (h *Handlers) ImagePage(c echo.Context) error {
	id := c.Param("id")

	image_inspect, _, err := h.docker_client.ImageInspectWithRaw(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	history_data, err := h.docker_client.ImageHistory(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	fmt.Printf("%+v\n", image_inspect)

	var history []imageHistoryModel
	for _, h := range history_data {
		history = append(history, make_image_history_model(h))
	}

	response := imagePageModel{
		Image:   make_image_model(image_inspect),
		History: history,
	}

	// fmt.Printf("%+v\n", response)

	return c.Render(http.StatusOK, "image", response)
}

func make_image_history_model(data image.HistoryResponseItem) imageHistoryModel {
	return imageHistoryModel{
		ID:      data.ID,
		Created: data.Created,
		Size:    data.Size,
		Tags:    data.Tags,
	}
}

func make_image_model(data types.ImageInspect) imageModel {
	return imageModel{
		ID:       data.ID,
		Size:     data.Size,
		RepoTags: data.RepoTags,
		Created:  data.Created,
		Parent:   data.Parent,
		Comment:  data.Comment,
	}
}

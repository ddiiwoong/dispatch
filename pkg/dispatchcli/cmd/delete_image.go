///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"golang.org/x/net/context"

	"github.com/spf13/cobra"

	"github.com/vmware/dispatch/pkg/api/v1"
	"github.com/vmware/dispatch/pkg/client"
	"github.com/vmware/dispatch/pkg/dispatchcli/i18n"
)

var (
	deleteImagesLong = i18n.T(`Delete images.`)

	// TODO: add examples
	deleteImagesExample = i18n.T(``)
)

// NewCmdDeleteImage creates command responsible for deleting  images.
func NewCmdDeleteImage(out io.Writer, errOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "image IMAGE_NAME",
		Short:   i18n.T("Delete images"),
		Long:    deleteImagesLong,
		Example: deleteImagesExample,
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"images"},
		Run: func(cmd *cobra.Command, args []string) {
			c := imageManagerClient()
			err := deleteImage(out, errOut, cmd, args, c)
			CheckErr(err)
		},
	}
	cmd.Flags().StringVarP(&cmdFlagApplication, "application", "a", "", "filter by application")
	return cmd
}

// CallDeleteImage makes the API call to delete an image
func CallDeleteImage(c client.ImagesClient) ModelAction {
	return func(i interface{}) error {
		imageModel := i.(*v1.Image)

		deleted, err := c.DeleteImage(context.TODO(), dispatchConfig.Organization, *imageModel.Name)
		if err != nil {
			return err
		}
		*imageModel = *deleted
		return nil
	}
}

func deleteImage(out, errOut io.Writer, cmd *cobra.Command, args []string, c client.ImagesClient) error {
	imageModel := v1.Image{
		Name: &args[0],
	}
	err := CallDeleteImage(c)(&imageModel)
	if err != nil {
		return err
	}
	return formatDeleteImageOutput(out, false, []*v1.Image{&imageModel})
}

func formatDeleteImageOutput(out io.Writer, list bool, images []*v1.Image) error {
	if dispatchConfig.JSON {
		encoder := json.NewEncoder(out)
		encoder.SetIndent("", "    ")
		if list {
			return encoder.Encode(images)
		}
		return encoder.Encode(images[0])
	}
	for _, i := range images {
		_, err := fmt.Fprintf(out, "Deleted image: %s\n", *i.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

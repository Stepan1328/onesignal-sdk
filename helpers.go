package onesignal

func NewCreateNotificationConfig() *CreateNotificationConfig {
	return &CreateNotificationConfig{}
}

func NewCancelNotificationConfig(id string) *CancelNotificationConfig {
	return &CancelNotificationConfig{
		NotificationID: id,
	}
}

func (c *CreateNotificationConfig) AddAndroidChannelID(channelID string) *CreateNotificationConfig {
	c.AndroidChannelID = channelID
	return c
}

func (c *CreateNotificationConfig) AddMessages(contents ...Content) *CreateNotificationConfig {
	for _, content := range contents {
		if c.Contents == nil {
			c.Contents = NewLocalizedContent()
		}
		c.Contents.Add(content.Location, content.Message)

		if c.Headings == nil {
			c.Headings = NewLocalizedContent()
		}
		c.Headings.Add(content.Location, content.Title)
	}

	return c
}

func (c *CreateNotificationConfig) AddDevices(ids ...string) *CreateNotificationConfig {
	if c.Devices == nil {
		c.Devices = NewDevices()
	}

	for _, id := range ids {
		c.Devices.AddDevice(id)
	}

	return c
}

func (c *CreateNotificationConfig) AddExternalUserIDs(ids ...string) *CreateNotificationConfig {
	if c.IncludeExternalUserIDs == nil {
		c.IncludeExternalUserIDs = NewDevices()
	}

	for _, id := range ids {
		c.IncludeExternalUserIDs.AddDevice(id)
	}

	return c
}

func (c *CreateNotificationConfig) AddData(data interface{}) *CreateNotificationConfig {
	c.Data = data

	return c
}

func (c *CreateNotificationConfig) AddAndroidGrouping(location, message string) *CreateNotificationConfig {
	if c.AndroidGroup == nil {
		c.AndroidGroup = NewLocalizedContent()
	}
	c.AndroidGroup.Add(location, message)

	return c
}

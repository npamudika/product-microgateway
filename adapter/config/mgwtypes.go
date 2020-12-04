/*
 *  Copyright (c) 2020, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package config

import (
	"sync"
	"time"
)

// Experimenting asynchronous communication between go routines using channels
// This uses singleton pattern where creating a single channel for communication
//
// To get a instance of the channel for a data publisher go routine
//  `publisher := NewSender()`
//
// Create a receiver channel in worker go routine
// receiver := NewReceiver()
//
// From publisher go routine, feed string value to the channel
// publisher<- "some value"
//
// In worker go routine, read the value sent by the publisher
// message := <-receiver
var once sync.Once

// C represents the channel to identify modifications added to the configuration file
// TODO: (VirajSalaka) remove this as unused.
var (
	C chan string // better to be interface{} type which could send any type of data.
)

// NewSender initializes the channel if it is not created an returns
func NewSender() chan string {
	once.Do(func() {
		C = make(chan string)
	})
	return C
}

// NewReceiver initializes the channel if it is not created an returns
func NewReceiver() chan string {
	once.Do(func() {
		C = make(chan string)
	})
	return C
}

// Config represents the adapter configuration.
// It is created directly from the configuration toml file.
type Config struct {

	// Server represents the configuration related to rest API (to which the apictl requests)
	Server struct {
		// Host name of the server
		Host string
		// Port of the server
		Port string
		// Public Certificate Path (For the https connection between adapter and apictl)
		PublicKeyPath string
		// Private Key Path (For the https connection between adapter and apictl)
		PrivateKeyPath string
		// Username credential
		Username string
		// Password credential
		Password string
	}

	// Envoy Listener Component related configurations.
	Envoy struct {
		ListenerHost            string
		ListenerPort            uint32
		ClusterTimeoutInSeconds time.Duration
		ListenerCertPath        string
		ListenerKeyPath         string
		ListenerTLSEnabled      bool
	}
}
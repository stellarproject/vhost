/*
	Copyright (c) 2019 Stellar Project

	Permission is hereby granted, free of charge, to any person
	obtaining a copy of this software and associated documentation
	files (the "Software"), to deal in the Software without
	restriction, including without limitation the rights to use, copy,
	modify, merge, publish, distribute, sublicense, and/or sell copies
	of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be
	included in all copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
	EXPRESS OR IMPLIED,
	INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
	IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
	HOLDERS BE LIABLE FOR ANY CLAIM,
	DAMAGES OR OTHER LIABILITY,
	WHETHER IN AN ACTION OF CONTRACT,
	TORT OR OTHERWISE,
	ARISING FROM, OUT OF OR IN CONNECTION WITH
	THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package util

import (
	tv1 "github.com/stellarproject/terraos/api/terra/v1"
	v1 "github.com/stellarproject/terraos/api/v1/orbit"
	"google.golang.org/grpc"
)

type LocalAgent struct {
	v1.AgentClient
	conn *grpc.ClientConn
}

func (a *LocalAgent) Close() error {
	return a.conn.Close()
}

func Agent(address string) (*LocalAgent, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &LocalAgent{
		AgentClient: v1.NewAgentClient(conn),
		conn:        conn,
	}, nil
}

type LocalTerra struct {
	tv1.TerraClient
	conn *grpc.ClientConn
}

func (l *LocalTerra) Close() error {
	return l.conn.Close()
}

func Terra(address string) (*LocalTerra, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &LocalTerra{
		TerraClient: tv1.NewTerraClient(conn),
		conn:        conn,
	}, nil
}

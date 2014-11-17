// Copyright (c) 2014 - Max Persson <max@looplab.se>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/maxpersson/otis2/master/eventhorizon/domain"
)

type CreateInvite struct {
	InvitationID domain.UUID
	Name         string
}

func (c CreateInvite) AggregateID() domain.UUID {
	return c.InvitationID
}

type AcceptInvite struct {
	InvitationID domain.UUID
}

func (c AcceptInvite) AggregateID() domain.UUID {
	return c.InvitationID
}

type DeclineInvite struct {
	InvitationID domain.UUID
}

func (c DeclineInvite) AggregateID() domain.UUID {
	return c.InvitationID
}

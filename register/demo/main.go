/* Copyright 2018 Bruce Liu.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"time"
	"goservice/register"
)

func main() {
	endpoints := []string{"119.23.67.84:2379", "120.79.38.97:2379"}
	register.Register("TestService2", "192.168.0.2", 8080, endpoints, 10*time.Second, 11)
	//register.Register()
	//register.UnRegister()

	time.Sleep(time.Second * 5000)
}

// -*- coding: utf-8 -*-

//  Copyright 2019 Pier Alberto <pieralbertopierini@gmail.com>
//
//  This program is free software; you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation; either version 3 of the License, or
//  (at your option) any later version.
//
//  This program is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU General Public License for more details.

//  You should have received a copy of the GNU General Public License
//  along with this program; if not, write to the Free Software
//  Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
//  MA 02110-1301, USA.
//
//

// Built on Linux

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/rdegges/go-ipify"
)

func main() {

	//Definitions of variables
	var oldIP string

	// Check presence of data file iplist.data if is not present create it and Open file
	datacsv, err := os.OpenFile("iplist.csv", os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		log.Println(err)
	}
	defer datacsv.Close()
	//Put the last IP data on the oldIP variable
	listcsv := csv.NewReader(datacsv)

	for {
		record, err := listcsv.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		oldIP = record[0]
	}
	fmt.Println("The old IP was: ", oldIP)
	//Check current IP
	newIP, err := ipify.GetIp()
	if err != nil {
		fmt.Println("Couldn't get my IP address:", err)
	} else {
		fmt.Println("My IP address is:", newIP)
	}
	// Check if old and new IP are equal
	if oldIP == newIP {
		fmt.Println("The IPs are equal")
	} else {
		fmt.Println("The IPs are different")
		//Append the new IP and data on the iplist.csv

		writeNewIP := csv.NewWriter(datacsv)

		//send an email and affix the new IP on the iplist.data if there are a new IP
	}

	//wait 5 minutes and restart the loop
}

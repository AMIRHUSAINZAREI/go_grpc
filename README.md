# Grpc in golang

This project is a playground to review grpc usage in golang

## Table of Contents

- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)

## Introduction

This project consist of two application. Calculator and Chat.

## Prerequisites

- Go (version 1.20)

## Installation

1. Clone the repository:

- git clone https://github.com/AMIRHUSAINZAREI/go_grpc_sample.git


2. Install dependencies:

- go mod download


## Usage

To Use Calculator applicatioon run a build a caculator server and run it, then build a calculator client and run it.
To Use Chat applicatioon run a build a chat server and run it, then build a chat client and run it.

## Configuration

To run the project, you need to create a `.env` file in the root directory and set the values for the following environment variables:

- `CALC_GRPC_SERVER_PORT`: port witch calculator server listen on.
- `CHAT_GRPC_SERVER_PORT`: Port witch chat server listen on.



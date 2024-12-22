# Tixara - Event Ticket Management System

  

Tixara is a simple Go-based application for managing events and ticket bookings. It enables users to add events, book tickets, view bookings and delete events with an easy-to-use command-line interface.

  

## Features

- Add new events with a specified number of tickets.

- Book tickets for an event and validate user input.

- Display all bookings for each event.

- Delete events when no longer needed.

- Sends ticket confirmation asynchronously.

  

## Prerequisites

- Go 1.20 or later installed on your system.

  

## Installation

1. Clone the repository:

```bash

git clone https://github.com/Vijey-Anbarasan/Tixara-Event-Ticket-Manager.git

```

2. Navigate to the project directory:

```bash

cd Tixara-Event-Ticket-Manager

```

  

## Usage

Run the application:

```bash

go  run  tixara.go  helper.go

```

  

Follow the on-screen menu to add events, book tickets, display bookings or delete events.

  

## Project Structure

-  `tixara.go`: Main application logic, including menu and event operations.

-  `helper.go`: Contains helper functions for input validation.

  

## Example Workflow

1. Add a new event:

- Enter the event name and the total number of tickets.

2. Book tickets for an event:

- Provide your name, email and the number of tickets to book.

- Receive a confirmation message asynchronously.

3. Display all bookings:

- View the list of all attendees for each event.

4. Delete an event:

- Remove an event when it’s no longer needed.

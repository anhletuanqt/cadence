# Assing To Chef

- Cadence/Camunda Task:
  - Wait for `preparation.assign_to_chef` signal from tenant admin/branch admin
- API Handler:
  - tenant admin/branch admin will call a api to choose a chef to assign the order
  - public message `preparation.assign_to_chef` to NATS

# Notify Chef

- Cadence/Camunda Task:
  - NATS req/reply to `messaging.notification.request` to the chef

# Start preparing

- Cadence/Camunda Task:
  - Wait for `preparation.new` signal from the chef
- API Handler:
  - the chef will call a api to add estimate time, preparing_at time and change the status of the order
  - public message `preparation.new` to NATS

# Notify Deliver:

- I don't know what you are using it for, but in my opinion we should remove it

# Update LCD

- Cadence/Camunda Task:
  - Public mesage `preparation.started` to NATS

# Set Order Is Ready

- Cadence/Camunda Task:
  - Wait for `preparation.ready-to-collect` signal from the chef
- API Handler:
  - the chef will call a api to add order_ready_at time and change the status of the order
  - public message `preparation.ready-to-collect` to NATS

# Notify Runner

- Cadence/Camunda Task:
  - NATS req/reply to `messaging.notification.request` to the runner (How will we identify that runner ??)

# L2 Consolidate Order

- Cadence/Camunda Task:
  - Wait for `preparation.in-consolidation` signal from the runner
- API Handler:
  - the runner will call a api to add the order at consolidation table
  - public message `preparation.in-consolidation` to NATS

# Put To Locker:

- Cadence/Camunda Task:
  - Wait for `preparation.in-locker` signal from the runner
- API Handler:
  - the runner will call a api to add in_locker time and change the status of the order
  - public message `preparation.in-locker` to NATS

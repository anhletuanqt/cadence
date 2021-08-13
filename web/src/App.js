import {
  Tabs,
  AppBar,
  Button,
  MenuItem,
  Typography,
  FormControl,
  InputLabel,
  Select,
  Card,
  CardContent,
  TextField,
} from '@material-ui/core';
import React from 'react';
import axios from 'axios';
import './App.css';

class App extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      value: 0,
      orders: [],
      statusInput: '',
      toggleStatus: false,
      tenantID: '',
      startPreparingOrder: '',
      preparingOrder: '',
      order: {
        tenantID: '',
        customerID: '',
      },
    };
  }

  loadOrders = async () => {
    const orders = await axios.get('http://localhost:8081/orders');
    this.setState({
      ...this.state,
      orders: orders.data.orders,
    });
  };

  async componentDidMount() {
    await this.loadOrders();
  }

  handleChangeStatus = async (event) => {
    const { tenantID } = this.state;
    const status = event.target.value;
    this.setState({
      ...this.state,
      statusInput: status,
    });

    let url = `http://localhost:8081/orders?status=${status}`;
    if (tenantID !== '') {
      url = `${url}&tenantID=${tenantID}`;
    }
    console.log(url);
    const orders = await axios.get(url);
    this.setState({
      ...this.state,
      orders: orders.data.orders,
    });
  };

  handleChangeTenantID = async (event) => {
    const { statusInput } = this.state;
    const id = event.target.value;
    // const { tenantID } = this.state;
    // const status = event.target.value;
    this.setState({
      ...this.state,
      tenantID: id,
    });
    let url = `http://localhost:8081/orders?tenantID=${id}`;
    if (statusInput !== '') {
      url = `${url}&status=${statusInput}`;
    }
    const orders = await axios.get(url);
    this.setState({
      ...this.state,
      orders: orders.data.orders,
    });
  };

  handleChangeStartPreparing = async (event) => {
    const id = event.target.value;
    // const { tenantID } = this.state;
    // const status = event.target.value;
    this.setState({
      ...this.state,
      startPreparingOrder: id,
    });
  };

  handleChangePreparing = async (event) => {
    const id = event.target.value;
    // const { tenantID } = this.state;
    // const status = event.target.value;
    this.setState({
      ...this.state,
      preparingOrder: id,
    });
  };

  handleCloseStatus = () => {
    this.setState({
      ...this.state,
      toggleStatus: false,
    });
  };

  handleOpenStatus = () => {
    this.setState({
      ...this.state,
      toggleStatus: true,
    });
  };

  clickPreparingOrder = async () => {
    const { preparingOrder } = this.state;
    await axios.post(
      `http://localhost:8081/orders/${preparingOrder}/preparing_order`
    );
    this.setState({
      ...this.state,
      preparingOrder: '',
    });

    const to = setTimeout(async () => {
      await this.loadOrders();
    }, 2000);
    clearTimeout(to);
  };

  clickStartPreparingOrder = async () => {
    const { startPreparingOrder } = this.state;
    await axios.post(
      `http://localhost:8081/orders/${startPreparingOrder}/start_preparing`
    );
    this.setState({
      ...this.state,
      startPreparingOrder: '',
    });

    const to = setTimeout(async () => {
      await this.loadOrders();
    }, 2000);
    clearTimeout(to);
  };

  handleChangeCreateOrderCustomer = (event) => {
    const id = event.target.value;
    const { order } = this.state;
    // const { tenantID } = this.state;
    // const status = event.target.value;
    this.setState({
      ...this.state,
      order: {
        ...order,
        customerID: id,
      },
    });
  };

  handleChangeCreateOrderTenant = (event) => {
    const id = event.target.value;
    const { order } = this.state;
    // const { tenantID } = this.state;
    // const status = event.target.value;
    this.setState({
      ...this.state,
      order: {
        ...order,
        tenantID: id,
      },
    });
  };

  clickCreateOrder = async () => {
    const { order } = this.state;
    if (order.customerID !== '' && order.tenantID !== '') {
      await axios.post('http://localhost:8081/orders', {
        tenantID: Number.parseInt(order.tenantID, 10),
        customerID: Number.parseInt(order.customerID, 10),
      });
      this.setState({
        ...this.state,
        order: {
          customerID: '',
          tenantID: '',
        },
      });
    }

    const to = setTimeout(async () => {
      await this.loadOrders();
    }, 2000);
    clearTimeout(to);
  };

  clickReload = async () => {
    this.loadOrders();
  };

  render() {
    const {
      orders,
      statusInput,
      tenantID,
      startPreparingOrder,
      preparingOrder,
      order,
    } = this.state;
    return (
      <div className="App" style={{ height: 800, paddingTop: 40 }}>
        <Card style={{ display: 'inline-block', width: '50%', height: '100%' }}>
          <FormControl style={{ width: 150 }}>
            <InputLabel id="demo-controlled-open-select-label">
              Status
            </InputLabel>
            <Select
              labelId="demo-controlled-open-select-label"
              id="demo-controlled-open-select"
              open={this.open}
              onClose={this.handleCloseStatus}
              onOpen={this.handleOpenStatus}
              value={statusInput}
              onChange={this.handleChangeStatus}
            >
              <MenuItem value={'pending'}>pending</MenuItem>
              <MenuItem value={'start preparing'}>start preparing</MenuItem>
              <MenuItem value={'preparing order'}>preparing order</MenuItem>
            </Select>
          </FormControl>
          <TextField
            id="orderID"
            label="TenantID"
            style={{ marginLeft: 20 }}
            value={tenantID}
            onChange={this.handleChangeTenantID}
          />
          <Button
            variant="contained"
            color="primary"
            style={{ marginLeft: 20, marginTop: 10 }}
            onClick={this.clickReload}
          >
            Reload
          </Button>
          {orders.map((order, index) => {
            return (
              <CardContent key={`card-c: ${index}`}>
                <Typography>ID: {order.ID}</Typography>
                <Typography>TenantID: {order.TenantID}</Typography>
                <Typography>CustomerID: {order.CustomerID}</Typography>
                <Typography>Status: {order.Status}</Typography>
                <Typography>
                  Activities:{' '}
                  {order.Activities.map((act, i) => {
                    return (
                      <Typography key={`p-c: ${i}`} style={{ margin: 0 }}>
                        {act}
                      </Typography>
                    );
                  })}
                </Typography>
              </CardContent>
            );
          })}
        </Card>
        <Card style={{ display: 'inline-block', width: '50%', height: '100%' }}>
          <Card style={{ height: '40%' }}>
            <Typography>Create Order</Typography>
            <TextField
              id="tenantID"
              label="TenantID"
              value={order.tenantID}
              onChange={this.handleChangeCreateOrderTenant}
            />
            <TextField
              id="customerID"
              label="CustomerID"
              style={{ marginLeft: 20 }}
              value={order.customerID}
              onChange={this.handleChangeCreateOrderCustomer}
            />
            <Button
              variant="contained"
              color="primary"
              style={{ marginLeft: 20, marginTop: 10 }}
              onClick={this.clickCreateOrder}
            >
              Create Order
            </Button>
          </Card>
          <Card style={{ height: '60%' }}>
            <Typography>Start Preparing Order</Typography>
            <TextField
              id="orderID"
              label="OrderID"
              style={{ marginLeft: 20 }}
              value={startPreparingOrder}
              onChange={this.handleChangeStartPreparing}
            />
            <Button
              variant="contained"
              color="primary"
              style={{ marginLeft: 10, marginTop: 10 }}
              onClick={this.clickStartPreparingOrder}
            >
              Start Preparing Order
            </Button>
            <Typography style={{ marginTop: 40 }}>Preparing Order</Typography>
            <TextField
              id="orderID"
              label="OrderID"
              style={{ marginLeft: 20 }}
              value={preparingOrder}
              onChange={this.handleChangePreparing}
            />
            <Button
              variant="contained"
              color="primary"
              style={{ marginLeft: 10, marginTop: 10 }}
              onClick={this.clickPreparingOrder}
            >
              Preparing Order
            </Button>
          </Card>
        </Card>
      </div>
    );
  }
}

export default App;

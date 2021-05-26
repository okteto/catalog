import React, { useState } from 'react';
import useInterval from 'use-interval';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import Divider from '@material-ui/core/Divider';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import Avatar from '@material-ui/core/Avatar';
import Typography from '@material-ui/core/Typography';

const POLLING_INTERVAL = 5000;

const useStyles = makeStyles(theme => ({
  root: {
    width: '100%',
    backgroundColor: theme.palette.background.paper,
    border: `1px solid ${theme.palette.grey[300]}`,
    borderLeft: 0,
    borderRight: 0
  },

  inline: {
    display: 'inline',
  }
}));

const fetchServices = async () => {
  try {
    const response = await fetch('/data');
    const data = await response.json();
    return Object.keys(data).map(id => {
      return { id, ...data[id] };
    });
  } catch(err) {
    console.error(err);
    return [];
  }
};

function Catalog() {
  const classes = useStyles();
  const [services, setServices] = useState([]);

  useInterval(async () => {
    const services = await fetchServices();
    if (services) {
      setServices(services);
    }
  }, POLLING_INTERVAL, true);

  return (
    <List className={classes.root}>
      {services.map((service, i) => (
        <div key={service.id}>
          <ListItem alignItems="flex-start">
            <ListItemText
              primary={service.service_name}
              secondary={
                <>
                  {service.owner_name}
                </>
              }
            />
          </ListItem>
          {i < services.length - 1 &&
            <Divider component="li" />
          }
        </div>
      ))}
    </List>
  );
}

export default Catalog;

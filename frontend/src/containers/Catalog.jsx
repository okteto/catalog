import React, { useState } from 'react';
import useInterval from 'use-interval';
import TimeAgo from 'javascript-time-ago';
import en from 'javascript-time-ago/locale/en';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import Divider from '@material-ui/core/Divider';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import Avatar from '@material-ui/core/Avatar';
import Typography from '@material-ui/core/Typography';

TimeAgo.addDefaultLocale(en);
const timeAgo = new TimeAgo('en-US');

const POLLING_INTERVAL = 5000;

const useStyles = makeStyles(theme => ({
  root: {
    width: '100%',
    backgroundColor: theme.palette.background.paper,
    border: `1px solid ${theme.palette.grey[300]}`,
    borderLeft: 0,
    borderRight: 0,
    padding: 0
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

const formatHealthTime = service => {
  const healthData = service.HealthData ?? [];
  if (healthData.length === 0) {
    return 'No data';
  }
  return timeAgo.format(Date.now() - Number(healthData[0].timestamp), 'round')
}

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
          <ListItem alignItems="flex-start" button>
            <ListItemText
              primary={service.service_name}
              secondary={
                <>
                  <Typography
                    component="span"
                    variant="body2"
                    className={classes.inline}
                    color="textPrimary"
                  >
                    {service.owner_name}
                  </Typography>
                  &nbsp;·&nbsp;
                  {formatHealthTime(service)}
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

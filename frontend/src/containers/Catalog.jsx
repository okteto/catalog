import React, { useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import Divider from '@material-ui/core/Divider';
import Avatar from '@material-ui/core/Avatar';
import Typography from '@material-ui/core/Typography';
import { deepPurple } from '@material-ui/core/colors';

import Health from './Health';

const useStyles = makeStyles(theme => ({
  root: {
    width: '100%',
    backgroundColor: theme.palette.background.paper,
    border: `1px solid ${theme.palette.grey[300]}`,
    borderLeft: 0,
    borderRight: 0,
    padding: 0
  },

  itemContent: {
    alignSelf: 'center',
    display: 'flex',
    flexDirection: 'column',
    flex: '1 auto',
    margin: theme.spacing(2)
  },

  owner: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center'
  },

  avatar: {
    color: theme.palette.getContrastText(deepPurple[500]),
    backgroundColor: deepPurple[500],
    marginRight: theme.spacing(0.5),
    height: '26px',
    width: '26px',
    fontSize: '1.2em'
  }
}));

function Catalog({ services = [] }) {
  const classes = useStyles();

  return (
    <List className={classes.root}>
      {services.map((service, i) => (
        <div key={service.id}>
          <ListItem alignItems="flex-start">
            <div className={classes.itemContent}>
              <Typography component="span" variant="h6" color="textPrimary" gutterBottom>
                {service.service_name}
              </Typography>
              <div className={classes.owner}>
                <Avatar className={classes.avatar}>
                  {service.owner_name[0]}
                </Avatar>
                <Typography component="span" variant="body1" color="textPrimary">
                  {service.owner_name}
                </Typography>
              </div>
            </div>
            <Health data={service.HealthData ?? []}/>
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

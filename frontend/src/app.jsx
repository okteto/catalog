import React from 'react';
import { render } from 'react-dom';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

import Catalog from './containers/Catalog';
import './app.css';

const useStyles = makeStyles(theme => ({
  header: {
    margin: '12px 24px'
  }
}));

function App() {
  const classes = useStyles();

  return (
    <>
      <Typography variant="h4" gutterBottom className={classes.header}>
        Service Catalog
      </Typography>
      <Catalog />
    </>
  );
}

render(<App />, document.getElementById('app'));

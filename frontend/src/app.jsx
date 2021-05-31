import React from 'react';
import { render } from 'react-dom';
import { makeStyles } from '@material-ui/core/styles';
import { createMuiTheme } from '@material-ui/core/styles';
import { ThemeProvider } from '@material-ui/styles';
import { blue } from '@material-ui/core/colors';
import Typography from '@material-ui/core/Typography';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';

import Catalog from './containers/Catalog';
import './app.css';

const theme = createMuiTheme({
  palette: {
    primary: {
      main: blue[700],
    },
    secondary: {
      main: '#11cb5f',
    }
  }
});

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },

  menuButton: {
    marginRight: theme.spacing(2),
  },

  title: {
    flexGrow: 1,
  }
}));

function App() {
  const classes = useStyles();

  return (
    <ThemeProvider theme={theme}>
      <div className={classes.root}>
        <AppBar position="static">
          <Toolbar>
            <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
              <MenuIcon />
            </IconButton>
            <Typography variant="h6" color="inherit" className={classes.title}>
              Service Catalog
            </Typography>
            <Button color="inherit">Refresh</Button>
          </Toolbar>
        </AppBar>
        <Catalog />
      </div>
    </ThemeProvider>
  );
}

render(<App />, document.getElementById('app'));

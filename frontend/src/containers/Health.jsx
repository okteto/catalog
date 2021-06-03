import React, { useState } from 'react';
import { format, fromUnixTime } from 'date-fns'
import { makeStyles } from '@material-ui/core/styles';
import { green, red } from '@material-ui/core/colors';
import CheckCircleIcon from '@material-ui/icons/CheckCircle';
import CancelIcon from '@material-ui/icons/Cancel';
import TimeAgo from 'javascript-time-ago';
import en from 'javascript-time-ago/locale/en';

TimeAgo.addDefaultLocale(en);
const timeAgo = new TimeAgo('en-US');

const formatHealthTime = timestamp => {
  if (!timestamp) {
    return 'Unknown';
  }
  return format(fromUnixTime(timestamp), 'MMM do KK:mm:ss');
}

const formatRelativeHealthTime = timestamp => {
  if (!timestamp || !Number(timestamp)) {
    return 'Unknown';
  }
  return timeAgo.format(Date.now() - Number(timestamp), 'round');
}

const useStyles = makeStyles((theme) => ({
  root: {
    alignItems: 'stretch',
    alignSelf: 'center',
    display: 'flex',
    flex: '1 auto',
    flexDirection: 'row',
    justifyContent: 'flex-start'
  },

  singleStatus: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    marginRight: theme.spacing(2)
  },

  timeline: {
    display: 'flex',
    flex: '1 auto',
    flexDirection: 'row',
    alignItems: 'center',
    marginRight: theme.spacing(2),
    maxWidth: '300px'
  },

  timelineItem: {
    position: 'relative',
    flexGrow: 1,
    borderBottom: '2px solid red',

    '&::after': {
      position: 'absolute',
      top: '-3px',
      right: 0,
      display: 'block',
      content: '""',
      width: '8px',
      height: '8px',
      borderRadius: '10px',
      backgroundColor: red[500]
    }
  },

  timelineStart: {
    marginRight: theme.spacing(1)
  },

  up: {
    borderColor: green[500],

    '&::after': {
      backgroundColor: green[500]
    }
  },

  down: {
    borderColor: red[500],

    '&::after': {
      backgroundColor: red[500]
    }
  },

  timelineEnd: {
    marginLeft: theme.spacing(1)
  }
}));

function Health({ data }) {
  const classes = useStyles();
  const [activeStep, setActiveStep] = useState(data.length);

  return (
    <div className={classes.root}>
      {data.length === 1 &&
        <div className={classes.singleStatus}>
          {data[0].healthy &&
            <CheckCircleIcon style={{ color: green[500] }} />
          }
          {!data[0].healthy &&
            <CancelIcon style={{ color: red[500] }} />
          }
        </div>
      }
      {data.length > 1 &&
        <div className={classes.timeline}>
          <div className={classes.timelineStart}>
            {formatRelativeHealthTime(data[data.length-1].timestamp)}
          </div>
          {data.reverse().map((item, i) => (
            <div
              className={`${classes.timelineItem} ${item.healthy ? classes.up : classes.down}`}
              key={i}
              title={formatHealthTime(item.timestamp)}
            />
          ))}
          <div className={classes.timelineEnd}>Now</div>
        </div>
      }
    </div>
  );
}

export default Health;
import React, { useState } from 'react';
import { format, fromUnixTime } from 'date-fns'
import { makeStyles } from '@material-ui/core/styles';
import { green, red } from '@material-ui/core/colors';
import Stepper from '@material-ui/core/Stepper';
import Step from '@material-ui/core/Step';
import StepLabel from '@material-ui/core/StepLabel';
import CheckCircleIcon from '@material-ui/icons/CheckCircle';
import CancelIcon from '@material-ui/icons/Cancel';

const formatHealthTime = timestamp => {
  if (!timestamp) {
    return 'Unknown';
  }
  return (
    <>
      {format(fromUnixTime(timestamp), 'MMM do')}
      <br />
      {format(fromUnixTime(timestamp), 'KK:mm:ss')}
    </>
  );
}

const useStyles = makeStyles((theme) => ({
  root: {
    alignItems: 'stretch',
    display: 'flex',
    flex: '1 auto',
    flexDirection: 'row',
    justifyContent: 'flex-end'
  },

  label: {
    '& .MuiStepLabel-alternativeLabel': {
      marginTop: theme.spacing(0.5),
      fontWeight: 'normal',
      fontSize: '.8rem',
      opacity: '.8'
    }
  },

  singleStatus: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    marginRight: theme.spacing(2)
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
        <Stepper activeStep={activeStep} alternativeLabel style={{ flex: '1' }}>
          {data.reverse().map((item, i) => (
            <Step key={i}>
              {item.healthy &&
                <StepLabel
                  classes={{ labelContainer: classes.label }}
                  icon={<CheckCircleIcon style={{ color: green[500] }} />}
                >
                  {formatHealthTime(item.timestamp)}
                </StepLabel>
              }
              {!item.healthy &&
                <StepLabel
                  classes={{ labelContainer: classes.label }}
                  icon={<CancelIcon style={{ color: red[500] }} />}
                >
                  {formatHealthTime(item.timestamp)}
                </StepLabel>
              }
            </Step>
          ))}
        </Stepper>
      }
    </div>
  );
}

export default Health;
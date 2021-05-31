import React, { useState } from 'react';
import TimeAgo from 'javascript-time-ago';
import en from 'javascript-time-ago/locale/en';
import { makeStyles } from '@material-ui/core/styles';
import { green, red } from '@material-ui/core/colors';
import Stepper from '@material-ui/core/Stepper';
import Step from '@material-ui/core/Step';
import StepLabel from '@material-ui/core/StepLabel';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import CheckCircleIcon from '@material-ui/icons/CheckCircle';
import CancelIcon from '@material-ui/icons/Cancel';

TimeAgo.addDefaultLocale(en);
const timeAgo = new TimeAgo('en-US');

const formatHealthTime = timestamp => {
  if (!timestamp) {
    return 'Unknown';
  }
  return timeAgo.format(Date.now() - Number(timestamp), 'round')
}

const useStyles = makeStyles((theme) => ({
  root: {},

  label: {
    '& .MuiStepLabel-alternativeLabel': {
      marginTop: theme.spacing(0.5),
      fontWeight: 'normal',
      opacity: '.8'
    }
  }
}));

function Health({ data }) {
  const classes = useStyles();
  const [activeStep, setActiveStep] = useState(data.length);

  return (
    <div className={classes.root}>
      <Stepper activeStep={activeStep} alternativeLabel>
        {data.map((item, i) => (
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
    </div>
  );
}

export default Health;
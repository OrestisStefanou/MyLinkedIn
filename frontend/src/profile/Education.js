import React from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import 'date-fns';

import DateFnsUtils from '@date-io/date-fns';
import {
  MuiPickersUtilsProvider,
  KeyboardDatePicker,
} from '@material-ui/pickers';
import { Container } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({

    form: {
      width: '100%', // Fix IE 11 issue.
      marginTop: theme.spacing(3),
    },
    submit: {
      margin: theme.spacing(3, 0, 2),
    },
  }));

export default function Education(){
  const classes = useStyles();

  const [selectedDate, setSelectedDate] = React.useState(new Date('2014-08-18T21:11:54'));

  const handleDateChange = (date) => {
    setSelectedDate(date);
  };

  return(
    <Container maxWidth="xs">
        <Typography component="h1" variant="h5">
        Add your education and experience
        </Typography>
        <form className={classes.form} noValidate>
            <Grid container spacing={2}>
            <Grid item xs={12} sm={6}>
                <TextField
                autoComplete="fname"
                name="degreeName"
                variant="outlined"
                required
                fullWidth
                id="degreeName"
                label="Degree Name"
                autoFocus
                />
            </Grid>
            <Grid item xs={12} sm={6}>
                <TextField
                variant="outlined"
                required
                fullWidth
                id="schoolName"
                label="School Name"
                name="schoolName"
                autoComplete="lname"
                />
            </Grid>
            <MuiPickersUtilsProvider utils={DateFnsUtils}>
            <Grid item xs={12} sm={6}>
            <KeyboardDatePicker
            disableToolbar
            variant="inline"
            format="MM/dd/yyyy"
            margin="normal"
            id="Entry Date"
            label="Date picker inline"
            value={selectedDate}
            onChange={handleDateChange}
            KeyboardButtonProps={{
                'aria-label': 'change date',
            }}
            />
            </Grid>
            <Grid item xs={12} sm={6}>
            <KeyboardDatePicker
            margin="normal"
            id="date-picker-dialog"
            label="Exit Date"
            format="MM/dd/yyyy"
            value={selectedDate}
            onChange={handleDateChange}
            KeyboardButtonProps={{
                'aria-label': 'change date',
            }}
            />
            </Grid>
        </MuiPickersUtilsProvider>
        </Grid>
        <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
        >
        Add
        </Button>
        </form>
    </Container>
  )
}
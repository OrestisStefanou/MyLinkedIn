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

export default function Experience(){
  const classes = useStyles();

  return(
    <Container maxWidth="xs">
        <Typography component="h1" variant="h5">
        Add your experience
        </Typography>
        <form className={classes.form} noValidate>
            <Grid container spacing={2}>
            <Grid item xs={12} sm={6}>
                <TextField
                autoComplete="fname"
                name="employeeName"
                variant="outlined"
                required
                fullWidth
                id="employeeName"
                label="Employee Name"
                autoFocus
                />
            </Grid>
            <Grid item xs={12} sm={6}>
                <TextField
                variant="outlined"
                required
                fullWidth
                id="jobTitle"
                label="Job Title"
                name="jobTitle"
                autoComplete="lname"
                />
            </Grid>
            <MuiPickersUtilsProvider utils={DateFnsUtils}>
            <Grid item xs={12} sm={6}>
            <TextField
            id="date"
            label="Start Date"
            type="date"
            className={classes.textField}
            InputLabelProps={{
            shrink: true,
            }}
            />
            </Grid>
            <Grid item xs={12} sm={6}>
            <TextField
            id="date"
            label="Finish Date"
            type="date"
            className={classes.textField}
            InputLabelProps={{
            shrink: true,
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
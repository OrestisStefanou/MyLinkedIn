import React, { useState,useEffect } from 'react';
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


import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import ListItemText from '@material-ui/core/ListItemText';
import Avatar from '@material-ui/core/Avatar';
import IconButton from '@material-ui/core/IconButton';
import SchoolIcon from '@material-ui/icons/School';
import DeleteIcon from '@material-ui/icons/Delete';

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
  const [educationInfo,setEducationInfo] = useState({degreeName:'',schoolName:'',startDate:'',finishDate:''});
  const [educationArray,setEducationArray] = useState([]);

  const handleChange = (e) => {
    const name = e.target.name;
    const value = e.target.value;
    setEducationInfo({ ...educationInfo, [name]: value });
  }; 

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(educationInfo);
    fetch('http://localhost:8080/v1/LinkedIn/addEducation', {
      method: "POST",
      mode:"cors",
      credentials:"include",
      body: JSON.stringify(educationInfo),
      headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
    })
    .then(response => response.json())
    .then((json) => {
      console.log(json);
      if(json.error){
        //Show error message
        console.log("SOMETHING WENT WRONG");
      }else{
        //Add the education info on the screen
        setEducationArray([...educationArray,educationInfo]);
        console.log("ALL GOOD");
      }
    });
  }

  useEffect(() => {
    fetch('http://localhost:8080/v1/LinkedIn/getEducation', {
      method: "GET",
      mode:"cors",
      credentials:"include",
      headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
    })
    .then(response => response.json())
    .then((json) => {
      if(json.error){
        //Show error message
        console.log(json.error);
      }else{
        //Add the education info on the screen
        console.log(json.education);
        setEducationArray(json.education);
      }
    });    
  },[]);

  return(
    <Container maxWidth="xs">
                <List dense={false}>
                  {educationArray.map((education) =>{
                    return(
                  <ListItem key={education.degreeName}>
                  <ListItemAvatar>
                    <Avatar>
                      <SchoolIcon />
                    </Avatar>
                  </ListItemAvatar>
                  <ListItemText
                    primary={education.degreeName + "  at  " + education.schoolName + "  From:  " + education.startDate + "  Until:  " + education.finishDate}
                  />
                  <ListItemSecondaryAction>
                    <IconButton edge="end" aria-label="delete">
                      <DeleteIcon />
                    </IconButton>
                  </ListItemSecondaryAction>
                </ListItem>
                    );
                  })}

                </List>

        <Typography component="h1" variant="h5">
        Add your education
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
                onChange={handleChange}
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
                onChange={handleChange}
                autoComplete="lname"
                />
            </Grid>
            <MuiPickersUtilsProvider utils={DateFnsUtils}>
            <Grid item xs={12} sm={6}>
            <TextField
            id="date"
            label="Start Date"
            type="date"
            name="startDate"
            onChange={handleChange}
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
            name="finishDate"
            onChange={handleChange}
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
            onClick={handleSubmit}
        >
        Add
        </Button>
        </form>
    </Container>
  )
}
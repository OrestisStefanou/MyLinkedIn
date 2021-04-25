import React, { useState,useEffect } from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import 'date-fns';
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

export default function Skills(){
  const classes = useStyles();
  const [skillInfo,setSkillInfo] = useState({name:''});
  const [skillsArray,setSkillsArray] = useState([]);

  const handleChange = (e) => {
    const name = e.target.name;
    const value = e.target.value;
    setSkillInfo({ ...skillInfo, [name]: value });
  }; 

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(skillInfo);
    fetch('http://localhost:8080/v1/LinkedIn/addSkill', {
      method: "POST",
      mode:"cors",
      credentials:"include",
      body: JSON.stringify(skillInfo),
      headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
    })
    .then(response => response.json())
    .then((json) => {
      if(json.error){
        //Show error message
        console.log(json.error);
      }else{
        //Add the education info on the screen
        setSkillsArray([...skillsArray,json.skillInfo]);
        console.log(json.skillInfo);
      }
    });
  }

  const removeSkill = (skillInfo) => {
    console.log("DELETETING ",skillInfo);
    fetch('http://localhost:8080/v1/LinkedIn/removeSkill', {
      method: "POST",
      mode:"cors",
      credentials:"include",
      body: JSON.stringify(skillInfo),
      headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
    })
    .then(response => response.json())
    .then((json) => {
      if(json.error){
        //Show error message
        console.log(json.error);
      }else{
        console.log(json.message);
        let newSkillsArray = skillsArray.filter((skill) => skill.name !== skillInfo.name);
        setSkillsArray(newSkillsArray);
      }
    });
  };

  useEffect(() => {
    fetch('http://localhost:8080/v1/LinkedIn/getSkills', {
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
        console.log(json.skills);
        if (json.skills!==null){
          setSkillsArray(json.skills);
        }
      }
    });    
  },[]);

  return(
    <Container maxWidth="xs">
                <List dense={false}>
                  {skillsArray && skillsArray.map((skill,index) =>{
                    return(
                  <ListItem key={index}>
                  <ListItemAvatar>
                    <Avatar>
                      <SchoolIcon />
                    </Avatar>
                  </ListItemAvatar>
                  <ListItemText
                    primary={skill.name }
                  />
                  <ListItemSecondaryAction>
                    <IconButton edge="end" aria-label="delete" onClick={() => removeSkill(skill)}>
                      <DeleteIcon />
                    </IconButton>
                  </ListItemSecondaryAction>
                </ListItem>
                    );
                  })}

                </List>

        <Typography component="h1" variant="h5">
        Add your Skills
        </Typography>
        <form className={classes.form} noValidate>
            <Grid container spacing={2}>
            <Grid item xs={12} sm={6}>
                <TextField
                autoComplete="fname"
                name="name"
                variant="outlined"
                required
                fullWidth
                id="name"
                label="Skill Name"
                onChange={handleChange}
                autoFocus
                />
            </Grid>
            <Grid item xs={12} sm={6}>
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
            </Grid>
        </Grid>
        </form>
    </Container>
  )
}
import React,{useEffect,useState} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import ListItemText from '@material-ui/core/ListItemText';
import Avatar from '@material-ui/core/Avatar';
import WorkIcon from '@material-ui/icons/Work';
import SchoolIcon from '@material-ui/icons/School';
import BuildIcon from '@material-ui/icons/Build';
import {useParams} from 'react-router-dom';

const useStyles = makeStyles({
  root: {
    maxWidth: 500,
  },
  media: {
    height: 200,
  },
});

export default function UserProfile() {
  const classes = useStyles();
  const [professionalInfo,setProfessionalInfo] = useState({});
  const [educationArray,setEducation] = useState([]);
  const [experienceArray,setExperience] = useState([]);
  const [skillsArray,setSkills] = useState([]);

  const professionalEmail = useParams().email;

  useEffect(() => {
    const getProfile = async () => {
        var url = 'http://localhost:8080/admin/LinkedIn/professional?email=' + professionalEmail; 
          fetch(url,{
              method: "GET",
              mode:"cors",
              credentials:"include",
              headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
              })
              .then(response => response.json())
              .then(json =>{
                  console.log(json);
                  setProfessionalInfo(json.professional);
                  if(json.education !== null){
                    setEducation(json.education);
                  }
                  if(json.experience !== null){
                    setExperience(json.experience);
                  }
                  if(json.skills !== null){
                    setSkills(json.skills);
                  }
              } )
              .catch(err => console.log('Request Failed',err))
        };
        getProfile();
  },[professionalEmail]);

  return (
    <Card className={classes.root}>
      <CardActionArea>
        <CardMedia
          className={classes.media}
          image={professionalInfo.photo}
          title="Profile picture"
        />
        <CardContent>
          <Typography gutterBottom variant="h5" component="h2">
            {professionalInfo.firstName + " " + professionalInfo.lastName}
          </Typography>
          <Typography gutterBottom variant="h5" component="h2">
            Email:{professionalInfo.email}
          </Typography>
          <Typography gutterBottom variant="h5" component="h2">
            Phone:{professionalInfo.phoneNumber}
          </Typography>
          <List dense={false}>
            {educationArray && educationArray.map((education,index) =>{
                return(
                  <ListItem key={index}>
                  <ListItemAvatar>
                    <Avatar>
                      <SchoolIcon />
                    </Avatar>
                  </ListItemAvatar>
                  <ListItemText
                    primary={education.degreeName + "  at  " + education.schoolName + "  From:  " + education.startDate + "  Until:  " + education.finishDate}
                  />
                </ListItem>
                );
            })}
        </List>

        <List dense={false}>
            {skillsArray && skillsArray.map((skill,index) =>{
                return(
                  <ListItem key={index}>
                  <ListItemAvatar>
                    <Avatar>
                      <BuildIcon />
                    </Avatar>
                  </ListItemAvatar>
                  <ListItemText
                    primary={skill.name }
                  />
                </ListItem>
                );
            })}
        </List>

        <List dense={false}>
            {experienceArray && experienceArray.map((experience,index) =>{
                return(
                  <ListItem key={index}>
                  <ListItemAvatar>
                    <Avatar>
                      <WorkIcon />
                    </Avatar>
                  </ListItemAvatar>
                  <ListItemText
                    primary={experience.jobTitle + "  at  " + experience.employerName + "  From:  " + experience.startDate + "  Until:  " + experience.finishDate}
                  />
                </ListItem>
                );
            })}
        </List>

        </CardContent>
      </CardActionArea>
      <CardActions>
      </CardActions>
    </Card>
  );
}
import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import ListItemText from '@material-ui/core/ListItemText';
import Avatar from '@material-ui/core/Avatar';
import WorkIcon from '@material-ui/icons/Work';
import SchoolIcon from '@material-ui/icons/School';
import BuildIcon from '@material-ui/icons/Build';


const useStyles = makeStyles({
  root: {
    maxWidth: 500,
  },
  media: {
    height: 200,
  },
});

export default function ProfileCard(props) {
  const classes = useStyles();
  const professionalInfo = props.professionalInfo;
  const educationArray = props.education;
  const experienceArray = props.experience;
  const skillsArray = props.skills;
  const selfProfile = props.selfProfile;

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
        {!selfProfile && 
        <Button size="small" color="primary">
          Send friend request
        </Button>
        }
      </CardActions>
    </Card>
  );
}
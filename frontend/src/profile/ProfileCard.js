import React,{useEffect,useState} from 'react';
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

  const [status,setStatus] = useState("");

  useEffect(() => {
    const getStatus = async () => {
      var url = 'http://localhost:8080/v1/LinkedIn/friendshipStatus?id=' + professionalInfo.id; 
        fetch(url,{
            method: "GET",
            mode:"cors",
            credentials:"include",
            headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
            })
            .then(response => response.json())
            .then(json =>{
                console.log(json.status);
                if(json.status.length > 0){
                  setStatus(json.status)
                }
            } )
            .catch(err => console.log('Request Failed',err))
      };
      getStatus();
  },[professionalInfo.id]);

  const sendFriendRequest = () => {
    const data = {professionalId2:professionalInfo.id}
    fetch('http://localhost:8080/v1/LinkedIn/addFriendRequest', {
      method: "POST",
      mode:"cors",
      credentials:"include",
      body: JSON.stringify(data),
      headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
      })
      .then(response => response.json())
          .then((json) => {
              if(json.error){
                  //Show error message
                  console.log(json.error);
              }else{
                  console.log(json);
                  setStatus("pending");
              }
        });       
  }

  const acceptFriendRequest = () => {
    const data = {professionalId2:professionalInfo.id}
    fetch('http://localhost:8080/v1/LinkedIn/addFriend', {
      method: "POST",
      mode:"cors",
      credentials:"include",
      body: JSON.stringify(data),
      headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
      })
      .then(response => response.json())
          .then((json) => {
              if(json.error){
                  //Show error message
                  console.log(json.error);
              }else{
                  console.log(json);
                  setStatus("friends");
              }
        });       
  }

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
        {!selfProfile && status.length === 0 &&
        <Button size="small" color="primary" onClick={sendFriendRequest}>
          Send friend request
        </Button>
        }
        {!selfProfile && status === "accept" &&
        <Button size="small" color="primary" onClick={acceptFriendRequest}>
          Accept Friend Request
        </Button>
        }
        {!selfProfile && status === "pending" &&
        <Typography variant="h6" gutterBottom>
          Friend request has been sent
        </Typography>
        }
        {!selfProfile && status === "friends" &&
        <Typography variant="h6" gutterBottom>
          Connected
        </Typography>
        }
      </CardActions>
    </Card>
  );
}
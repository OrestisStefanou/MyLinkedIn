import React,{useState,useEffect} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardHeader from '@material-ui/core/CardHeader';
import CardMedia from '@material-ui/core/CardMedia';
import CardContent from '@material-ui/core/CardContent';
import Avatar from '@material-ui/core/Avatar';
import { red } from '@material-ui/core/colors';
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import CloudUploadIcon from '@material-ui/icons/CloudUpload';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';

const useStyles = makeStyles((theme) => ({
  root: {
    maxWidth: 500,
  },
  media: {
    height: 0,
    paddingTop: '56.25%', // 16:9
  },
  expand: {
    transform: 'rotate(0deg)',
    marginLeft: 'auto',
    transition: theme.transitions.create('transform', {
      duration: theme.transitions.duration.shortest,
    }),
  },
  expandOpen: {
    transform: 'rotate(180deg)',
  },
  avatar: {
    backgroundColor: red[500],
  },
}));

export default function RecipeReviewCard() {
  const classes = useStyles();
  const [profileInfo,setProfileInfo] = useState({});

  useEffect(() => {
    fetch('http://localhost:8080/v1/LinkedIn/authenticated', {
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
        //Set profile Info
        console.log(json.professional);
        setProfileInfo(json.professional);
      }
    });    
  },[]);

  return (
    <Card className={classes.root}>
      <CardHeader
        avatar={
          <Avatar aria-label="recipe" className={classes.avatar}>
            <AccountCircleIcon/>
          </Avatar>
        }
        title="My Profile"
      />
      <CardMedia
        className={classes.media}
        image={profileInfo.photo}
        title="Paella dish"
      />
      <CardContent>
      <form className={classes.form} noValidate>
          <Grid container spacing={2}>
            <Grid item xs={12} sm={6}>
              <TextField
                name="firstName"
                variant="outlined"
                required
                fullWidth
                id="firstName"
                label="First Name"
                defaultValue={profileInfo.firstName}
                placeholder={profileInfo.firstName}
                autoFocus
              />
            </Grid>
            <Grid item xs={12} sm={6}>
              <TextField
                variant="outlined"
                required
                fullWidth
                id="lastName"
                label="Last Name"
                defaultValue={profileInfo.lastName}
                placeholder={profileInfo.lastName}
                name="lastName"
                autoComplete="lname"
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                variant="outlined"
                required
                fullWidth
                id="email"
                label="Email"
                defaultValue={profileInfo.email}
                placeholder={profileInfo.email}
                name="email"
                autoComplete="email"
              />
            </Grid>
            <Grid item xs={12} sm={6}>
              <TextField
                variant="outlined"
                required
                fullWidth
                name="password"
                label="New Password"
                type="password"
                id="password"
                autoComplete="current-password"
              />
            </Grid>
            <Grid item xs={12} sm={6}>
              <TextField
                variant="outlined"
                required
                fullWidth
                name="password2"
                label="Confirm Password"
                type="password"
                id="password"
                autoComplete="current-password"
              />
            </Grid>
            <Grid item xs={12} sm={6}>
              <TextField
                variant="outlined"
                required
                fullWidth
                id="lastName"
                label="Phone Number"
                defaultValue={profileInfo.phoneNumber}
                placeholder={profileInfo.phoneNumber}
                name="phone_number"
                autoComplete="phone number"
              />
            </Grid>
            <Grid item xs={12} sm={6}>
            <Button
                    variant="contained"
                    className={classes.button}
                    color="default"
                    component="label"
                    startIcon={<CloudUploadIcon />}
                  >Change Profile Picture
                <input type="file" id="image" name="image" hidden  />
                </Button>
            </Grid>
          </Grid>
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
          >
            Update Profile
          </Button>
        </form>
      </CardContent>
    </Card>
  );
}

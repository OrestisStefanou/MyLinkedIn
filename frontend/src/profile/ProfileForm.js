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
import Alert from '@material-ui/lab/Alert';


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
  const [profileInfo,setProfileInfo] = useState({email:'',firstName:'',lastName:'',password:'',password2:'',phoneNumber:''});
  const [image, setImage] = useState({});
  const [photoPicked,setPhotoPicked] = useState(false);
  const [errorMessage,setErrorMessage] = useState('');

  const handleChange = (e) => {
    const name = e.target.name;
    const value = e.target.value;
    setProfileInfo({ ...profileInfo, [name]: value });
  };

  const handleImageChange = (e) => {
    setImage(e.target.files[0]);
    setPhotoPicked(true);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(profileInfo);
    console.log(image);
    //If new password given
    if(profileInfo.password.length !== 0){
      //Check if passwords match
      if (profileInfo.password !== profileInfo.password2){
        setErrorMessage("Passwords do not match");
        return;
      }
      //Check if password is at least 8 characters
      if(profileInfo.password.length < 8){
        setErrorMessage("Passwords must have at least 8 characters");
        return;  
      }
    }
    if(profileInfo.firstName.length === 0){
      setErrorMessage("First Name can not be empty");
    }
    if(profileInfo.lastName.length === 0){
      setErrorMessage("Last Name can not be empty");
    }
    if(profileInfo.email.length === 0){
      setErrorMessage("Email can not be empty");
    }
    let form_data = new FormData();
    form_data.append('email',profileInfo.email);
    form_data.append('firstName',profileInfo.firstName);
    form_data.append('lastName',profileInfo.lastName);
    form_data.append('password',profileInfo.password);
    form_data.append('phoneNumber',profileInfo.phoneNumber);
    if (photoPicked){
      form_data.append('photo',image,image.name);
    }
    fetch(
      'http://localhost:8080/v1/LinkedIn/updateProfessional',
      {
        method: 'POST',
        mode:"cors",
        credentials:"include",
        body: form_data,
      }
    )
    .then((response) => response.json())
    .then((result) => {
      console.log(result);
      if(result.error){
        setErrorMessage(result.error);
      }else{
        //Show success message
        //history.push(`/signin`)
      }
    })
  };

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
        image={profileInfo.photo ? profileInfo.photo : 'www.exaple.com'}
        title="Paella dish"
      />
      <CardContent>
      {errorMessage && <Alert onClose={() => setErrorMessage("")} severity="error">{errorMessage}</Alert>}
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
                onChange={handleChange}
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
                onChange={handleChange}
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
                onChange={handleChange}
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
                onChange={handleChange}
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
                onChange={handleChange}
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
                onChange={handleChange}
                name="phoneNumber"
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
                <input type="file" id="image" name="image" hidden onChange={handleImageChange} />
                </Button>
                <p>{image.name}</p>
            </Grid>
          </Grid>
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onClick={handleSubmit}
          >
            Update Profile
          </Button>
        </form>
      </CardContent>
    </Card>
  );
}

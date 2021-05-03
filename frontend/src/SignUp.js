import React,{useState,useEffect} from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import Alert from '@material-ui/lab/Alert';
import CloudUploadIcon from '@material-ui/icons/CloudUpload';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import { useHistory } from 'react-router-dom';


function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        LinkedIn
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(3),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
  appBar: {
    borderBottom: `1px solid ${theme.palette.divider}`,
    background : '#3f51b5',
    color: "#FFFFFF",
  },
  toolbar: {
    flexWrap: 'wrap',
  },
  toolbarTitle: {
    flexGrow: 1,
  },
}));

export default function SignUp() {
  const classes = useStyles();
  const [image, setImage] = useState({});
  const [photoPicked,setPhotoPicked] = useState(false);
  const [userInfo,setUserInfo] = useState({email:'',first_name:'',last_name:'',password:'',password2:'',phone_number:''});
  const [errorMessage,setErrorMessage] = useState('');

  let history = useHistory();

  const handleChange = (e) => {
      const name = e.target.name;
      const value = e.target.value;
      setUserInfo({ ...userInfo, [name]: value });
  };

  const handleImageChange = (e) => {
    setImage(e.target.files[0]);
    setPhotoPicked(true);
  };

  const handleSubmit = (e) => {
      e.preventDefault();
      console.log(userInfo);
      console.log(image);

      //Check if passwords match
      if (userInfo.password !== userInfo.password2){
        setErrorMessage("Passwords do not match");
        return;
      }
      //Check if password is at least 8 characters
      if(userInfo.password.length < 8){
        setErrorMessage("Passwords must have at least 8 characters");
        return;  
      }
      let form_data = new FormData();
      form_data.append('email',userInfo.email);
      form_data.append('firstName',userInfo.first_name);
      form_data.append('lastName',userInfo.last_name);
      form_data.append('password',userInfo.password);
      form_data.append('phoneNumber',userInfo.phone_number);
      if (photoPicked){
        form_data.append('photo',image,image.name);
      }
      fetch(
        'http://localhost:8080/v1/LinkedIn/signup',
        {
          method: 'POST',
          mode:"cors",
          credentials:"include",
          body: form_data,
        }
      )
      .then((response) => response.json())
      .then((result) => {
        if(result.error){
          setErrorMessage(result.error);
        }else{
          history.push(`/signin`);
        }
      })
  };

  useEffect(() => {
    const checkSession = async () => {
      const response = await fetch('http://localhost:8080/v1/LinkedIn/authenticated',{
        method: "GET",
        mode:"cors",
        credentials:"include",
        headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
        });
      const jsonResponse = await response.json();
      console.log(jsonResponse);
      if (response.status === 202) {
        history.push(`/home`);
      }
    };
    checkSession();
  },[history]);

  return (
    <React.Fragment>
      <CssBaseline />
      <AppBar position="static" color="default" elevation={0} className={classes.appBar}>
        <Toolbar className={classes.toolbar}>
          <Typography variant="h4" color="inherit" noWrap className={classes.toolbarTitle}>
          <Link color="inherit" href="/">LinkedIn</Link>
          </Typography>
        </Toolbar>
      </AppBar>
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign up
        </Typography>
        {errorMessage && <Alert onClose={() => setErrorMessage("")} severity="error">{errorMessage}</Alert>}
        <form className={classes.form} noValidate>
          <Grid container spacing={2}>
            <Grid item xs={12} sm={6}>
              <TextField
                autoComplete="fname"
                name="first_name"
                variant="outlined"
                required
                fullWidth
                id="firstName"
                label="First Name"
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
                name="last_name"
                autoComplete="lname"
                onChange={handleChange}
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                variant="outlined"
                required
                fullWidth
                id="email"
                label="Email Address"
                name="email"
                autoComplete="email"
                onChange={handleChange}
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                variant="outlined"
                required
                fullWidth
                name="password"
                label="Password"
                type="password"
                id="password"
                autoComplete="current-password"
                onChange={handleChange}
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                variant="outlined"
                required
                fullWidth
                name="password2"
                label="Confirm Password"
                type="password"
                id="password"
                autoComplete="current-password"
                onChange={handleChange}
              />
            </Grid>
            <Grid item xs={12} sm={6}>
              <TextField
                variant="outlined"
                required
                fullWidth
                id="lastName"
                label="Phone number"
                name="phone_number"
                autoComplete="phone number"
                onChange={handleChange}
              />
            </Grid>
            <Grid item xs={12} sm={6}>
            <Button
                    variant="contained"
                    className={classes.button}
                    color="default"
                    component="label"
                    startIcon={<CloudUploadIcon />}
                  >Upload an image
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
            Sign Up
          </Button>
          <Grid container justify="flex-end">
            <Grid item>
              <Link href="/signin" variant="body2">
                Already have an account? Sign in
              </Link>
            </Grid>
          </Grid>
        </form>
      </div>
      <Box mt={5}>
        <Copyright />
      </Box>
    </Container>
    </React.Fragment>
  );
}
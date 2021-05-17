import React, { useState,useEffect } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import Alert from '@material-ui/lab/Alert';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import { useHistory } from 'react-router-dom';
import SupervisedUserCircleIcon from '@material-ui/icons/SupervisedUserCircle';


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
    marginTop: theme.spacing(1),
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

export default function AdminSignIn() {
  const classes = useStyles();
  const [adminLoginInfo,setAdminLoginInfo] = useState({email:'',password:''});
  const [errorMessage,setErrorMessage] = useState('');

  let history = useHistory();

  const handleChange = (e) => {
    const name = e.target.name;
    const value = e.target.value;
    setAdminLoginInfo({ ...adminLoginInfo, [name]: value });
  }; 

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(adminLoginInfo);
    fetch('http://localhost:8080/admin/LinkedIn/signin', {
      method: "POST",
      mode:"cors",
      credentials:"include",
      body: JSON.stringify(adminLoginInfo),
      headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
    })
    .then(response => response.json())
    .then((json) => {
      console.log(json);
      if(json.error){
        setErrorMessage(json.error);
      }else{
        history.push(`/admin/home`);
      }
    });
  }

  useEffect(() => {
    const checkSession = async () => {
      const response = await fetch('http://localhost:8080/admin/LinkedIn/authenticated',{
        method: "GET",
        mode:"cors",
        credentials:"include",
        headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
        });
      const jsonResponse = await response.json()
      console.log(jsonResponse)
      if (response.status === 202) {
        //history.push(`/home`);  //Chang this to the homepage of the admin
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
          <SupervisedUserCircleIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Admin Sign in
        </Typography>
        {errorMessage && <Alert onClose={() => setErrorMessage("")} severity="error">{errorMessage}</Alert>}
        <form className={classes.form} noValidate>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            onChange={handleChange}
            autoFocus
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            onChange={handleChange}
          />
          <FormControlLabel
            control={<Checkbox value="remember" color="primary" />}
            label="Remember me"
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onClick={handleSubmit}
          >
            Sign In
          </Button>
          <Grid container>
          </Grid>
        </form>
      </div>
      <Box mt={8}>
        <Copyright />
      </Box>
    </Container>
    </React.Fragment>
  );
}
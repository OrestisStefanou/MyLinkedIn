import React from 'react';
import CssBaseline from '@material-ui/core/CssBaseline';
import Link from '@material-ui/core/Link';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import 'date-fns';
import ProfileForm from "./ProfileForm";
import Education from "./Education";
import Experience from "./Experience";

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

export default function Profile() {
  const classes = useStyles();

  return (
    <React.Fragment>
      <CssBaseline />
      <AppBar position="static" color="default" elevation={0} className={classes.appBar}>
        <Toolbar className={classes.toolbar}>
          <Typography variant="h4" color="inherit" noWrap className={classes.toolbarTitle}>
          <Link color="inherit" href="/home">LinkedIn</Link>
          </Typography>
        </Toolbar>
      </AppBar>
    <Container component="main" maxWidth="md">
      <CssBaseline />
      <div className={classes.paper}>
        <ProfileForm/>
      </div>
      <div className={classes.paper}>
        <Education/>
      </div>
      <div className={classes.paper}>
        <Experience/>
      </div>
      <Box mt={5}>
        <Copyright />
      </Box>
    </Container>
    </React.Fragment>
  );
}
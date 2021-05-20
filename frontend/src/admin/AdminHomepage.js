import React,{useState} from 'react';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Box from '@material-ui/core/Box';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Link from '@material-ui/core/Link';
import Button from '@material-ui/core/Button';
import { useHistory } from 'react-router-dom';
import { useEffect } from 'react';
import Title from "../homepage/Title";
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import DraftsIcon from '@material-ui/icons/Drafts';
import GroupIcon from '@material-ui/icons/Group';

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

const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
  },
  toolbar: {
    paddingRight: 24, // keep right padding when drawer closed
  },
  toolbarIcon: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'flex-end',
    padding: '0 8px',
    ...theme.mixins.toolbar,
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },
  appBarShift: {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  menuButton: {
    marginRight: 36,
  },
  menuButtonHidden: {
    display: 'none',
  },
  title: {
    flexGrow: 1,
  },
  drawerPaper: {
    position: 'relative',
    whiteSpace: 'nowrap',
    width: drawerWidth,
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  drawerPaperClose: {
    overflowX: 'hidden',
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    width: theme.spacing(7),
    [theme.breakpoints.up('sm')]: {
      width: theme.spacing(9),
    },
  },
  appBarSpacer: theme.mixins.toolbar,
  content: {
    flexGrow: 1,
    height: '100vh',
    overflow: 'auto',
  },
  container: {
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  },
  paper: {
    padding: theme.spacing(2),
    display: 'flex',
    overflow: 'auto',
    flexDirection: 'column',
  },
  fixedHeight: {
    height: 240,
  },
}));

export default function AdminHomepage() {
  const classes = useStyles();
  const [adminInfo,setAdminInfo] = useState({});

  const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight);
  let history = useHistory();

  const handleLogout = () => {
    fetch('http://localhost:8080/v1/LinkedIn/logout',{
    method: "GET",
    mode:"cors",
    credentials:"include",
    headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
    })
    .then(response => response.json())
    .then(json => console.log(json))
    .catch(err => console.log('Request Failed',err))
    history.push("/");
  }

  useEffect(() => {
    const checkSession = async () => {
        const response = await fetch('http://localhost:8080/admin/LinkedIn/authenticated',{
          method: "GET",
          mode:"cors",
          credentials:"include",
          headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
          });
        const jsonResponse = await response.json();
        console.log(jsonResponse);
        if (response.status !== 202) {
          console.log("Something went wrong");
          history.push("/");
        }else{
          console.log(jsonResponse.admin);
          setAdminInfo(jsonResponse.admin);
        }
      };
      checkSession();
  },[history]);

  return (
    <div className={classes.root}>
      <CssBaseline />
      <AppBar position="absolute" className={clsx(classes.appBar, false && classes.appBarShift)}>
        <Toolbar className={classes.toolbar}>
          <Typography component="h1" variant="h6" color="inherit" noWrap className={classes.title}>
            <Link color="inherit" href="/admin/home">LinkedIn</Link>
          </Typography>
          <Button variant="contained" color="secondary" onClick={handleLogout}>
            Logout
          </Button>
        </Toolbar>
      </AppBar>
      <main className={classes.content}>
        <div className={classes.appBarSpacer} />
        <Container maxWidth="lg" className={classes.container}>
          <Grid container spacing={3}>
            <Grid item xs={12} md={8} lg={9}>
              <Paper className={classes.paper}>
                <List component="nav" aria-label="main mailbox folders">
                    <ListItem button onClick={() => history.push(`/admin/users`)}>
                    <ListItemIcon>
                        <GroupIcon />
                    </ListItemIcon>
                    <ListItemText primary="See the users" />
                    </ListItem>
                    <ListItem button>
                    <ListItemIcon>
                        <DraftsIcon />
                    </ListItemIcon>
                    <ListItemText primary="Drafts" />
                    </ListItem>
                </List>
              </Paper>
            </Grid>
            {/* Admin Info */}
            <Grid item xs={12} md={4} lg={3}>
              <Paper className={fixedHeightPaper}>
              <React.Fragment>
                <Title>Admin user</Title>  
                <Title>{adminInfo.firstName + " " + adminInfo.lastName}</Title>           
                <Title> Email:{adminInfo.email} </Title>
                <Title> Phone Number:{adminInfo.phoneNumber}</Title>
              </React.Fragment>
              </Paper>
            </Grid>
          </Grid>
          <Box pt={4}>
            <Copyright />
          </Box>
        </Container>
      </main>
    </div>
  );
}
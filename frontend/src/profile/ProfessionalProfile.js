import React,{useState} from 'react';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Drawer from '@material-ui/core/Drawer';
import Box from '@material-ui/core/Box';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import IconButton from '@material-ui/core/IconButton';
import Badge from '@material-ui/core/Badge';
import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Link from '@material-ui/core/Link';
import MenuIcon from '@material-ui/icons/Menu';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import NotificationsIcon from '@material-ui/icons/Notifications';
import ListItems  from '../homepage/ListItems';
import Deposits from '../homepage/Deposits';
import Button from '@material-ui/core/Button';
import { useHistory } from 'react-router-dom';
import { useEffect } from 'react';
import ChatIcon from '@material-ui/icons/Chat';
import {useParams} from 'react-router-dom';
import ProfileCard from "./ProfileCard";

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

export default function ProfessionalProfile() {
  const classes = useStyles();
  const [open, setOpen] = useState(true);
  const [professionalInfo,setProfessionalInfo] = useState({});
  const [education,setEducation] = useState([]);
  const [experience,setExperience] = useState([]);
  const [skills,setSkills] = useState([]);

  const handleDrawerOpen = () => {
    setOpen(true);
  };
  const handleDrawerClose = () => {
    setOpen(false);
  };

  const handleShowNotifications = () => {
    history.push(`/notifications`);
  }

  const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight);
  let history = useHistory();
  const professionalEmail = useParams().email;
  let selfProfile = false;
  if (professionalEmail === "self"){
    selfProfile = true;
  }

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
    const getProfile = async () => {
      var url = 'http://localhost:8080/v1/LinkedIn/professional?email=' + professionalEmail; 
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
    <div className={classes.root}>
      <CssBaseline />
      <AppBar position="absolute" className={clsx(classes.appBar, open && classes.appBarShift)}>
        <Toolbar className={classes.toolbar}>
          <IconButton
            edge="start"
            color="inherit"
            aria-label="open drawer"
            onClick={handleDrawerOpen}
            className={clsx(classes.menuButton, open && classes.menuButtonHidden)}
          >
            <MenuIcon />
          </IconButton>
          <Typography component="h1" variant="h6" color="inherit" noWrap className={classes.title}>
            <Link color="inherit" href="/home">LinkedIn</Link>
          </Typography>
          <IconButton color="inherit">
            <Badge badgeContent={0} color="secondary" onClick={()=>history.push(`/messages`)}>
              <ChatIcon />
            </Badge>
          </IconButton>
          <IconButton color="inherit">
            <Badge badgeContent={0} color="secondary" onClick={handleShowNotifications}>
              <NotificationsIcon />
            </Badge>
          </IconButton>
          <Button variant="contained" color="secondary" onClick={handleLogout}>
            Logout
          </Button>
        </Toolbar>
      </AppBar>
      <Drawer
        variant="permanent"
        classes={{
          paper: clsx(classes.drawerPaper, !open && classes.drawerPaperClose),
        }}
        open={open}
      >
        <div className={classes.toolbarIcon}>
          <IconButton onClick={handleDrawerClose}>
            <ChevronLeftIcon />
          </IconButton>
        </div>
        <Divider />
        <ListItems/>
        <Divider />
      </Drawer>
      <main className={classes.content}>
        <div className={classes.appBarSpacer} />
        <Container maxWidth="lg" className={classes.container}>
          <Grid container spacing={3}>
            <Grid item xs={12} md={8} lg={9}>
                <ProfileCard professionalInfo={professionalInfo} education={education} experience={experience} skills={skills} selfProfile={selfProfile}/>
            </Grid>
            {/* Recent Deposits */}
            <Grid item xs={12} md={4} lg={3}>
              <Paper className={fixedHeightPaper}>
                <Deposits />
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
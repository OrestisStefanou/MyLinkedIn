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
import ListItems  from './ListItems';
import Deposits from './Deposits';
import Button from '@material-ui/core/Button';
import { useHistory } from 'react-router-dom';
import { useEffect } from 'react';
import ChatIcon from '@material-ui/icons/Chat';
import TextField from '@material-ui/core/TextField';
import ProfessionalsList from "./ProfessionalsList"

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

export default function NetworkPage() {
  const classes = useStyles();
  const [open, setOpen] = useState(false);
  const [searchResults,setSearchResults] = useState([]);
  const [connectedProfessionals,setConnectedProfessionals] = useState([]);

  const handleChange = (e) => {
    const value = e.target.value;
    if (value.length === 0){
      setSearchResults([]);
      return;
    }
    fetch('http://localhost:8080/v1/LinkedIn/searchProfessional?'+ new URLSearchParams({
        query: value,
    }),{
        method: "GET",
        mode:"cors",
        credentials:"include",
        headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
        })
        .then(response => response.json())
        .then(json =>{
          if (json.results !== null){
            setSearchResults(json.results);
          }
        } )
        .catch(err => console.log('Request Failed',err))
  }; 
  
  const handleDrawerOpen = () => {
    setOpen(true);
  };
  const handleDrawerClose = () => {
    setOpen(false);
  };

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
    const getConnectedProfessionals = async () => {
      const response = await fetch('http://localhost:8080/v1/LinkedIn/connectedProfessionals',{
        method: "GET",
        mode:"cors",
        credentials:"include",
        headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
        });
      const jsonResponse = await response.json();
      console.log(jsonResponse);
      if (response.status !== 200) {
        history.push(`/`);
      }else{
        console.log(jsonResponse.connectedProfessionals);
        if(jsonResponse.connectedProfessionals !== null){
          setConnectedProfessionals(jsonResponse.connectedProfessionals);
        }
      }
    };
    getConnectedProfessionals();
  },[history]);

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
            <Badge badgeContent={0} color="secondary" >
              <ChatIcon />
            </Badge>
          </IconButton>
          <IconButton color="inherit">
            <Badge badgeContent={0} color="secondary" onClick={()=>history.push(`/notifications`)}>
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
          <TextField
            id="standard-full-width"
            label="Search a professional"
            style={{ margin: 8 }}
            placeholder="Search by first name,last name or email"
            fullWidth
            margin="normal"
            InputLabelProps={{
                shrink: true,
            }}
            onChange={handleChange}
          />
            <Grid item xs={12} md={8} lg={9}>
            {searchResults.length > 0 && 
              <Paper className={fixedHeightPaper}>
                <Typography component="h6" variant="h6" color="inherit" noWrap className={classes.title}>
                    <ProfessionalsList professionals={searchResults}/>
                </Typography>
              </Paper>
            }
            </Grid>
            {/* Recent Deposits */}
            <Grid item xs={12} md={4} lg={3}>
              <Paper className={fixedHeightPaper}>
                <Deposits />
              </Paper>
            </Grid>
            <Grid item xs={12}>
              <Paper className={classes.paper}>
                <Typography component="h6" variant="h6" color="inherit" noWrap className={classes.title}>
                  Connected Professionals with me
                </Typography>
                <ProfessionalsList professionals={connectedProfessionals}/>
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
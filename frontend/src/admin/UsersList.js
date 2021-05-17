import React,{useState} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import Divider from '@material-ui/core/Divider';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import Avatar from '@material-ui/core/Avatar';
import Typography from '@material-ui/core/Typography';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import Checkbox from '@material-ui/core/Checkbox';
import { Button } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100%',
    maxWidth: '40ch',
    backgroundColor: theme.palette.background.paper,
  },
  inline: {
    display: 'inline',
  },
}));

export default function UsersList(props) {
  const classes = useStyles();
  const [selectedUsersID,setSelectedUsersID] = useState([]);

  const selectUser = (user) => {
    if (user.checked === false){
        user.checked = true;
        setSelectedUsersID([...selectedUsersID,user.userInfo.id]);
    }else{
        user.checked = false;
        let newUsersIDArray = selectedUsersID.filter((id) => id !== user.userInfo.id );
        setSelectedUsersID(newUsersIDArray);
    }
  }

  return (
    <React.Fragment>
    <Button variant="outlined" color="primary" onClick={() => console.log(selectedUsersID)}>
    Export to json
    </Button>
    <Button variant="outlined" color="default">
    Export to XML
    </Button>
    <List className={classes.root}>
        {props.users.map((user) => {
            return(
                <React.Fragment>
                    <ListItem alignItems="flex-start">
                    <ListItemIcon>
                    <Checkbox
                        onClick={() => selectUser(user)}
                        edge="start"
                        checked={user.checked}
                        tabIndex={-1}
                        disableRipple
                        inputProps={{ 'aria-labelledby': user.userInfo.id }}
                    />
                    </ListItemIcon>
                        <ListItemAvatar>
                            <Avatar alt="Profile Photo" src={user.userInfo.photo} />
                        </ListItemAvatar>
                        <ListItemText
                            primary={user.userInfo.firstName + " " + user.userInfo.lastName}
                            secondary={
                                <React.Fragment>
                                <Typography
                                    component="span"
                                    variant="body2"
                                    className={classes.inline}
                                    color="textPrimary"
                                >
                                    {user.userInfo.email}
                                </Typography>
                                <Button variant="outlined" color="primary">
                                    Show profile
                                </Button>
                                </React.Fragment>
                            }
                        />
                    </ListItem>
                    
                    <Divider variant="inset" component="li" />
                </React.Fragment>
            )
        })}
    </List>
    </React.Fragment>
  );
}
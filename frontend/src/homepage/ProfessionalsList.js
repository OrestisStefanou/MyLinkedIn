import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import Divider from '@material-ui/core/Divider';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import Avatar from '@material-ui/core/Avatar';

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100%',
    maxWidth: '36ch',
    backgroundColor: theme.palette.background.paper,
  },
  inline: {
    display: 'inline',
  },
}));

export default function ProfessionalsList(props) {
  const classes = useStyles();

  return (
    <List className={classes.root}>
        {props.professionals.map((professional) => {
            return(
                <React.Fragment>
                <ListItem button onClick={() => {console.log(professional)}} alignItems="flex-start">
                <ListItemAvatar>
                  <Avatar alt="Profile picture" src={professional.photo} />
                </ListItemAvatar>
                <ListItemText
                  primary={professional.firstName + " " + professional.lastName}
                />
              </ListItem>
              <Divider variant="inset" component="li" />
              </React.Fragment>
            )
        })}
    </List>
  );
}
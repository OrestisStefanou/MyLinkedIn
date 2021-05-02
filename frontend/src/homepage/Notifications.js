import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100%',
    maxWidth: 360,
    backgroundColor: theme.palette.background.paper,
    position: 'relative',
    overflow: 'auto',
    maxHeight: 300,
  },
  listSection: {
    backgroundColor: 'inherit',
  },
  ul: {
    backgroundColor: 'inherit',
    padding: 0,
  },
}));

export default function Notifications(props) {
  const classes = useStyles();

  return (
    <List className={classes.root} subheader={<li />}>
          <ul className={classes.ul}>
            {props.notifications.map((notification) => (
              <ListItem key={notification.id}>
                <ListItemText primary={notification.msg} />
              </ListItem>
            ))}
          </ul>
    </List>
  );
}
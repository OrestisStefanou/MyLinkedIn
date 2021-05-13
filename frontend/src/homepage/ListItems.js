import React from 'react';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import PeopleIcon from '@material-ui/icons/People';
import PersonIcon from '@material-ui/icons/Person';
import DescriptionIcon from '@material-ui/icons/Description';
import SettingsIcon from '@material-ui/icons/Settings';
import List from '@material-ui/core/List';
import { useHistory } from 'react-router-dom';

export default function ListItems() {
  let history = useHistory();
  return(
  <List>
  <div>
    <ListItem button onClick={()=>history.push(`/professionals/self`)}>
      <ListItemIcon>
        <PersonIcon />
      </ListItemIcon>
      <ListItemText primary="My profile" />
    </ListItem>
    <ListItem button onClick={()=>history.push(`/jobs`)}>
      <ListItemIcon>
        <DescriptionIcon />
      </ListItemIcon>
      <ListItemText primary="Job ads" />
    </ListItem>
    <ListItem button onClick={()=>history.push(`/network`)}>
      <ListItemIcon>
        <PeopleIcon />
      </ListItemIcon>
      <ListItemText primary="My Network" />
    </ListItem>
    <ListItem button onClick={()=>history.push(`/settings`)}>
      <ListItemIcon>
        <SettingsIcon/>
      </ListItemIcon>
      <ListItemText primary="Settings" />
    </ListItem>
  </div>
  </List>
  )
};
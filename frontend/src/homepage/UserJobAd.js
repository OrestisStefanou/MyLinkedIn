import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import clsx from 'clsx';
import Card from '@material-ui/core/Card';
import CardHeader from '@material-ui/core/CardHeader';
import CardContent from '@material-ui/core/CardContent';
import CardActions from '@material-ui/core/CardActions';
import Collapse from '@material-ui/core/Collapse';
import Avatar from '@material-ui/core/Avatar';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';
import { red } from '@material-ui/core/colors';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import MoreVertIcon from '@material-ui/icons/MoreVert';
import ProfessionalList from "./ProfessionalsList";
import DeleteIcon from '@material-ui/icons/Delete';
import AttachFileIcon from '@material-ui/icons/AttachFile';
import Alert from '@material-ui/lab/Alert';

const useStyles = makeStyles((theme) => ({
  root: {
    maxWidth: 545,
  },
  media: {
    height: 0,
    paddingTop: '56.25%', // 16:9
  },
  expand: {
    transform: 'rotate(0deg)',
    marginLeft: 'auto',
    transition: theme.transitions.create('transform', {
      duration: theme.transitions.duration.shortest,
    }),
  },
  expandOpen: {
    transform: 'rotate(180deg)',
  },
  avatar: {
    backgroundColor: red[500],
  },
}));

export default function RecipeReviewCard(props) {
  const classes = useStyles();
  const [expanded, setExpanded] = React.useState(false);
  const [message,setMessage] = React.useState("");

  const adInfo = props.adInfo.ad;
  const interestedProfessionals = props.adInfo.interestedProfessionals

  const handleExpandClick = () => {
    setExpanded(!expanded);
  };

  const removeJobAd = (adInfo) => {
    console.log(adInfo);
    fetch('http://localhost:8080/v1/LinkedIn/removeJobAd', {
      method: "POST",
      mode:"cors",
      credentials:"include",
      body: JSON.stringify(adInfo),
      headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
    })
    .then(response => response.json())
    .then((json) => {
      if(json.error){
        //Show error message
        console.log(json.error);
      }else{
        console.log(json.message);
        setMessage("Job Ad deleted.Please refresh the page");
      }
    });
  }

  return (
    <Card className={classes.root}>
      {message && <Alert onClose={() => setMessage("")} severity="success">{message}</Alert>}
      <CardHeader
        avatar={
          <Avatar aria-label="recipe" className={classes.avatar}>
            {adInfo.id}
          </Avatar>
        }
        action={
          <IconButton aria-label="settings">
            <MoreVertIcon />
          </IconButton>
        }
        title={adInfo.title}
        subheader=""
      />
      <CardContent>
        <Typography variant="body2" color="textSecondary" component="p">
            {adInfo.jobDescription}
        </Typography>
      </CardContent>
      <CardActions disableSpacing>
        <IconButton aria-label="add to favorites" onClick={() => removeJobAd(adInfo)}>
          <DeleteIcon />
        </IconButton>
        { adInfo.file &&
        <IconButton
          className={clsx(classes.expand, {
            [classes.expandOpen]: expanded,
          })}
          aria-expanded={expanded}
          href={adInfo.file}
          aria-label="get file"
        >
          <AttachFileIcon />
        </IconButton>
        }
        <IconButton
          className={clsx(classes.expand, {
            [classes.expandOpen]: expanded,
          })}
          onClick={handleExpandClick}
          aria-expanded={expanded}
          aria-label="show more"
        >
          <ExpandMoreIcon />
        </IconButton>
        <Typography>Interested Professionals</Typography>
      </CardActions>
      <Collapse in={expanded} timeout="auto" unmountOnExit>
        <CardContent>
          {interestedProfessionals && <ProfessionalList professionals={interestedProfessionals}/>}
        </CardContent>
      </Collapse>
    </Card>
  );
}
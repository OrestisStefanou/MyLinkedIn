import React,{useState,useEffect} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import clsx from 'clsx';
import Card from '@material-ui/core/Card';
import CardHeader from '@material-ui/core/CardHeader';
import CardMedia from '@material-ui/core/CardMedia';
import CardContent from '@material-ui/core/CardContent';
import CardActions from '@material-ui/core/CardActions';
import Collapse from '@material-ui/core/Collapse';
import Avatar from '@material-ui/core/Avatar';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';
import { red } from '@material-ui/core/colors';
import MoreVertIcon from '@material-ui/icons/MoreVert';
import ChatBubbleOutlineOutlinedIcon from '@material-ui/icons/ChatBubbleOutlineOutlined';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import AttachFileIcon from '@material-ui/icons/AttachFile';
import GradeIcon from '@material-ui/icons/Grade';
import GradeOutlinedIcon from '@material-ui/icons/GradeOutlined';

const useStyles = makeStyles((theme) => ({
  root: {
    maxWidth: 700,
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
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

export default function JobAd(props) {
  const classes = useStyles();
  const [expanded, setExpanded] = React.useState(false);
  const [uploaderInfo,setUploaderInfo] = useState({firstName:"",lastName:""});
  const [hasImage,setHasImage] = useState(false);
  const [like,setLike] = useState(false);
  const [comment,setComment] = useState({adId:props.adInfo.id,comment:""});
  const [adComments,setAdComments] = useState([]);
  const [adInterest,setAdInterest] = useState(0);

  const handleExpandClick = () => {
    setExpanded(!expanded);
  };

  const handleChange = (e) => {
    const name = e.target.name;
    const value = e.target.value;
    setComment({ ...comment, [name]: value });
  };

  const handleCommentSubmit = (e) => {
    e.preventDefault();
    fetch('http://localhost:8080/v1/LinkedIn/article/addComment', {     //Change the endpoint here
      method: "POST",
      mode:"cors",
      credentials:"include",
      body: JSON.stringify(comment),
      headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
      })
      .then(response => response.json())
          .then((json) => {
              if(json.error){
                  //Show error message
                  console.log(json.error);
              }else{
                  //Show the comment on the comment section
                  console.log(json);
                  setAdComments([...adComments,json.comment]);
              }
        });       
  };

  const handleLike = () => {    //Change the name to handle interest
    if (like === false){
      fetch('http://localhost:8080/v1/LinkedIn/jobAd/addInterest', {  
        method: "POST",
        mode:"cors",
        credentials:"include",
        body: JSON.stringify(props.articleInfo),
        headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
        })
        .then(response => response.json())
            .then((json) => {
                if(json.error){
                    //Show error message
                    console.log(json.error);
                }else{
                    console.log(json);
                    setLike(true);
                    setAdInterest(adInterest+1);
                }
          });       
    }else{
      fetch('http://localhost:8080/v1/LinkedIn/jobAd/removeInterest', {   
        method: "POST",
        mode:"cors",
        credentials:"include",
        body: JSON.stringify(props.articleInfo),
        headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
        })
        .then(response => response.json())
            .then((json) => {
                if(json.error){
                    //Show error message
                    console.log(json.error);
                }else{
                    console.log(json);
                    setLike(false);
                    setAdInterest(adInterest-1);
                }
          });       
    }
  };

  useEffect(() => {
    fetch('http://localhost:8080/v1/LinkedIn/getJobAdDetails', {  //Change the endpoint 
    method: "POST",
    mode:"cors",
    credentials:"include",
    body: JSON.stringify(props.adInfo),
    headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
    })
    .then(response => response.json())
        .then((json) => {
            if(json.error){
                //Show error message
                console.log(json.error);
            }else{
                console.log(json);
                setUploaderInfo(json.uploader);
                setHasImage(json.hasImage);
                setLike(json.liked);
                if (json.comments !== null){
                  setAdComments(json.comments);
                }
                setAdInterest(json.likes);
            }
        });    
  },[props.adInfo]);

  return (
    <Card className={classes.root}>
      <CardHeader
        avatar={
          <Avatar aria-label="recipe" className={classes.avatar}>
            {uploaderInfo.firstName[0] + uploaderInfo.lastName[0]}
          </Avatar>
        }
        action={
          <IconButton aria-label="settings">
            <MoreVertIcon />
          </IconButton>
        }
        title={props.adInfo.title}
        subheader={uploaderInfo.firstName + " " + uploaderInfo.lastName}
      />
      { hasImage &&
      <CardMedia
        className={classes.media}
        image={props.adInfo.file}
        title="Image"
      />
        }
      <CardContent>
        <Typography variant="body2" color="textPrimary" component="p">
            {props.adInfo.jobDescription}
        </Typography>
      </CardContent>
      <CardActions disableSpacing>
        <IconButton aria-label="add to favorites" onClick={handleLike}>
        {like ? <GradeIcon/> : <GradeOutlinedIcon /> }
        </IconButton>
        {adInterest}
        <IconButton aria-label="share" onClick={handleExpandClick}>
          <ChatBubbleOutlineOutlinedIcon />
        </IconButton>
        {adComments.length}
        { !hasImage && props.adInfo.file &&
        <IconButton
          className={clsx(classes.expand, {
            [classes.expandOpen]: expanded,
          })}
          aria-expanded={expanded}
          href={props.adInfo.file}
          aria-label="get file"
        >
          <AttachFileIcon />
        </IconButton>
        }
      </CardActions>
      <Collapse in={expanded} timeout="auto" unmountOnExit>
        <CardContent>
        <form className={classes.form} noValidate>
          <Grid container spacing={1}>
            <Grid item xs={12} sm={8}>
              <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                name="comment"
                label="Comment"
                type="text"
                id="content"
                onChange={handleChange}
                autoComplete="content"
              />
            </Grid>
            <Grid item xs={12} sm={4}>
              <Button
                type="submit"
                fullWidth
                variant="contained"
                color="primary"
                className={classes.submit}
                onClick={handleCommentSubmit}
              >
                Comment
              </Button>
            </Grid>
          </Grid>
          </form>
          {/*Show the comments */}
          {adComments.map((adComment) => {
            return(
              <Typography key={adComment.id} variant="subtitle1" gutterBottom>
                {adComment.firstName + " " + adComment.lastName + ":" + adComment.comment}
              </Typography>
            )
          })}
        </CardContent>
      </Collapse>
    </Card>
  );
}
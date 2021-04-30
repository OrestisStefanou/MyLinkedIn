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
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import MoreVertIcon from '@material-ui/icons/MoreVert';
import ChatBubbleOutlineOutlinedIcon from '@material-ui/icons/ChatBubbleOutlineOutlined';
import ThumbUpAltOutlinedIcon from '@material-ui/icons/ThumbUpAltOutlined';
import ThumbUpAltIcon from '@material-ui/icons/ThumbUpAlt';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

const useStyles = makeStyles((theme) => ({
  root: {
    maxWidth: 345,
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

export default function Article(props) {
  const classes = useStyles();
  const [expanded, setExpanded] = React.useState(false);
  const [uploaderInfo,setUploaderInfo] = useState({firstName:"",lastName:""});
  const [hasImage,setHasImage] = useState(false);
  const [like,setLike] = useState(false);
  const [comment,setComment] = useState({articleId:props.articleInfo.id,comment:""});
  const [articleComments,setArticleComments] = useState([]);

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
    fetch('http://localhost:8080/v1/LinkedIn/article/addComment', {
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
                  setArticleComments([...articleComments,json.comment]);
              }
        });       
  };

  const handleLike = () => {
    if (like === false){
      fetch('http://localhost:8080/v1/LinkedIn/article/addLike', {
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
                }
          });       
    }else{
      fetch('http://localhost:8080/v1/LinkedIn/article/removeLike', {
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
                }
          });       
    }
  };

  useEffect(() => {
    fetch('http://localhost:8080/v1/LinkedIn/getArticleDetails', {
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
                setUploaderInfo(json.uploader);
                setHasImage(json.hasImage);
                setLike(json.liked);
                if (json.comments !== null){
                  setArticleComments(json.comments);
                }
            }
        });    
  },[props.articleInfo]);

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
        title={props.articleInfo.title}
        subheader={uploaderInfo.firstName + " " + uploaderInfo.lastName}
      />
      { hasImage &&
      <CardMedia
        className={classes.media}
        image={props.articleInfo.file}
        title="Paella dish"
      />
        }
      <CardContent>
        <Typography variant="body2" color="textPrimary" component="p">
            {props.articleInfo.content}
        </Typography>
      </CardContent>
      <CardActions disableSpacing>
        <IconButton aria-label="add to favorites" onClick={handleLike}>
        {like ? <ThumbUpAltIcon/> : <ThumbUpAltOutlinedIcon /> }
        </IconButton>
        <IconButton aria-label="share" onClick={handleExpandClick}>
          <ChatBubbleOutlineOutlinedIcon />
        </IconButton>
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
          {articleComments.map((articleComment) => {
            return(
              <Typography key={articleComment.id} variant="subtitle1" gutterBottom>
                {articleComment.firstName + " " + articleComment.lastName + ":" + articleComment.comment}
              </Typography>
            )
          })}
        </CardContent>
      </Collapse>
    </Card>
  );
}
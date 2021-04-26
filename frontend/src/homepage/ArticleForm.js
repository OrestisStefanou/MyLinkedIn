import React,{useState} from 'react';
import Title from './Title';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Container from '@material-ui/core/Container';
import AttachFileIcon from '@material-ui/icons/AttachFile';
import Grid from '@material-ui/core/Grid';
import Alert from '@material-ui/lab/Alert';


const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(1),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
  appBar: {
    borderBottom: `1px solid ${theme.palette.divider}`,
    background : '#3f51b5',
    color: "#FFFFFF",
  },
  toolbar: {
    flexWrap: 'wrap',
  },
  toolbarTitle: {
    flexGrow: 1,
  },
}));

export default function ArticleForm() {
  const classes = useStyles();
  const [file,setFile] = useState({});
  const [filePicked,setFilePicked] = useState(false);
  const [articleInfo,setArticleInfo] = useState({title:"",content:""});
  const [errorMessage,setErrorMessage] = useState('');
  const [successMessage,setSuccessMessage] = useState('');

  const handleChange = (e) => {
    const name = e.target.name;
    const value = e.target.value;
    setArticleInfo({ ...articleInfo, [name]: value });
  };

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
    setFilePicked(true);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(articleInfo);
    console.log(file);

    //Check if title given
    if(articleInfo.title.length === 0){
      setErrorMessage("Please enter a title");
      return;  
    }
    //Check if content given
    if(articleInfo.content.length === 0){
      setErrorMessage("Please enter a content");
      return;  
    }
    let form_data = new FormData();
    form_data.append('title',articleInfo.title);
    form_data.append('content',articleInfo.content);
    if (filePicked){
      form_data.append('file',file,file.name);
    }
    fetch(
      'http://localhost:8080/v1/LinkedIn/addArticle',
      {
        method: 'POST',
        mode:"cors",
        credentials:"include",
        body: form_data,
      }
    )
    .then((response) => response.json())
    .then((result) => {
      if(result.error){
        setErrorMessage(result.error);
      }else{
        setSuccessMessage('Article created');
        setArticleInfo({title:"",content:""});
        setFile({});
      }
    })
};

  return (
    <React.Fragment>
      <Container component="main" maxWidth="md">
      <Title>Add an article</Title>
      <CssBaseline />
      {errorMessage && <Alert onClose={() => setErrorMessage("")} severity="error">{errorMessage}</Alert>}
      <div className={classes.paper}>
        <form className={classes.form} noValidate>
          <Grid container spacing={1}>
            <Grid item xs={8}>
              <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                id="title"
                label="Article Title"
                name="title"
                autoComplete="articleTitle"
                onChange={handleChange}
                autoFocus
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                multiline
                rows={2}
                name="content"
                label="Content"
                type="text"
                id="content"
                autoComplete="content"
                onChange={handleChange}
              />
            </Grid>
            <Grid item xs={12} sm={6}>
              <Button
                variant="contained"
                fullWidth
                className={classes.submit}
                color="default"
                component="label"
                startIcon={<AttachFileIcon />}
              >Attach a file
              <input type="file" id="file" name="file" hidden onChange={handleFileChange}/>
              </Button>
              <p>{file.name}</p>
            </Grid>
            <Grid item xs={12} sm={6}>
              <Button
                type="submit"
                fullWidth
                variant="contained"
                color="primary"
                className={classes.submit}
                onClick={handleSubmit}
              >
                Post article
              </Button>
              {successMessage && <Alert onClose={() => setSuccessMessage("")} severity="success">{successMessage}</Alert>}
            </Grid>
          </Grid>
          </form>
        </div>
      
    </Container>
    </React.Fragment>
  );
}
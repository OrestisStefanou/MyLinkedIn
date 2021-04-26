import React from 'react';
import Title from './Title';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Container from '@material-ui/core/Container';
import AttachFileIcon from '@material-ui/icons/AttachFile';
import Grid from '@material-ui/core/Grid';

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

  return (
    <React.Fragment>
      <Container component="main" maxWidth="md">
      <Title>Add an article</Title>
      <CssBaseline />
      <div className={classes.paper}>
        <form className={classes.form} noValidate>
          <Grid container spacing={1}>
            <Grid item xs={8}>
              <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                id="email"
                label="Article Title"
                name="email"
                autoComplete="articleTitle"
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
                name="password"
                label="Content"
                type="text"
                id="password"
                autoComplete="current-password"
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
              <input type="file" id="image" name="image" hidden />
              </Button>
              <p></p>
            </Grid>
            <Grid item xs={12} sm={6}>
              <Button
                type="submit"
                fullWidth
                variant="contained"
                color="primary"
                className={classes.submit}
              >
                Post article
              </Button>
            </Grid>
          </Grid>
          </form>
        </div>
      
    </Container>
    </React.Fragment>
  );
}
import React from "react";
import { config } from "./config";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import { Grid, TextField, MenuItem, Button, Paper } from "@material-ui/core";
import Snackbar from "@material-ui/core/Snackbar";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      "& .MuiTextField-root": {
        margin: theme.spacing(1),
        width: 200
      }
    },
    paper: {
      marginTop: theme.spacing(3),
      marginBottom: theme.spacing(3),
      padding: theme.spacing(2),
      [theme.breakpoints.up(600 + theme.spacing(3) * 2)]: {
        marginTop: theme.spacing(6),
        marginBottom: theme.spacing(6),
        padding: theme.spacing(3)
      }
    }
  })
);

function Alert(props: AlertProps) {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
}

function Home() {
  const classes = useStyles();
  const [word, setWord] = React.useState("");
  const [category, setCategory] = React.useState("名");
  const [mean, setMean] = React.useState("");
  const [any, setAny] = React.useState("");
  const [submitting, setSubmitting] = React.useState(false);
  const [success, setSuccess] = React.useState(false);
  const [error, setError] = React.useState(false);

  const addWord = async (
    word: string,
    category: string,
    mean: string,
    any: string
  ) => {
    setSubmitting(true);

    fetch('http://localhost:1998/vocabulary', {
      method: 'POST',
      mode: 'no-cors',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ word: word, category: category, mean: mean, any: any })
    });

    setSubmitting(false);
    clearForm();
    setSuccess(true);

  };

  const clearForm = () => {
    setWord("");
    setCategory("名");
    setMean("");
    setAny("");
  };

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    switch (event.target.name) {
      case "word":
        setWord(event.target.value);
        break;
      case "category":
        setCategory(event.target.value);
        break;
      case "mean":
        setMean(event.target.value);
        break;
      case "any":
        setAny(event.target.value);
        break;
      default:
        alert("key not found");
    }
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    addWord(word, category, mean, any);
  };

  const handleClose = () => {
    if (success) {
      setSuccess(false);
    }
    if (error) {
      setError(false);
    }
  };

  return (
    <div className="App">
      <div>
        <Paper className={classes.paper}>
          <form
            className={classes.root}
            onSubmit={handleSubmit}
            autoComplete="off"
          >
            <Grid
              container
              direction="column"
              justify="space-evenly"
              alignItems="flex-start"
            >
              <Grid item>
                <TextField
                  id="outlined-basic"
                  name="word"
                  label={config.form.label.word}
                  variant="outlined"
                  onChange={handleChange}
                  value={word}
                  required
                />
              </Grid>

              <Grid item>
                <TextField
                  id="outlined-select-currency"
                  name="category"
                  select
                  label={config.form.label.category}
                  value={category}
                  onChange={handleChange}
                  variant="outlined"
                  required
                >
                  {config.categories.map(option => (
                    <MenuItem key={option.value} value={option.value}>
                      {option.label}
                    </MenuItem>
                  ))}
                </TextField>
              </Grid>

              <Grid item>
                <TextField
                  id="outlined-basic"
                  name="mean"
                  label={config.form.label.mean}
                  variant="outlined"
                  onChange={handleChange}
                  value={mean}
                  required
                />
              </Grid>

              <Grid item>
                <TextField
                  id="outlined-multiline-static"
                  name="any"
                  label={config.form.label.any}
                  multiline
                  rows="4"
                  variant="outlined"
                  onChange={handleChange}
                  value={any}
                />
              </Grid>

              <Grid item>
                <Button
                  disabled={submitting ? true : false}
                  type="submit"
                  size="medium"
                  variant="contained"
                  color="primary"
                >
                  {config.message.add}
                </Button>
              </Grid>
            </Grid>
          </form>

          <Snackbar
            open={success}
            autoHideDuration={6000}
            onClose={handleClose}
          >
            <Alert onClose={handleClose} severity="success">
              {config.message.success}
            </Alert>
          </Snackbar>

          <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="error">
              {config.message.error}
            </Alert>
          </Snackbar>
        </Paper>
      </div>
    </div>
  );
}

export default Home;

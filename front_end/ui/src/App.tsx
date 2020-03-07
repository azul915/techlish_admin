import React from "react";
import axios from "axios";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import { Grid, TextField, MenuItem, Button, Paper } from "@material-ui/core";
import "./styles.css";

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

const categories = [
  {
    value: "名",
    label: "名詞"
  },
  {
    value: "動",
    label: "動詞"
  },
  {
    value: "形",
    label: "形容詞"
  },
  {
    value: "副",
    label: "副詞"
  }
];

function App() {
  const addWord = async (
    word: string,
    category: string,
    mean: string,
    any: string
  ) => {
    try {
      const result = await axios.post(
        `http://localhost:1998/vocabulary?word=${word}&category=${category}&mean=${mean}&any=${any}`
      );
      console.log(result.data);
    } catch {
      console.log("error!!");
    }
  };

  const classes = useStyles();
  const [word, setWord] = React.useState("");
  const [category, setCategory] = React.useState("名");
  const [mean, setMean] = React.useState("");
  const [any, setAny] = React.useState("");
  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    switch (event.target.name) {
      case "word":
        console.log(event.target.value);
        setWord(event.target.value);
        break;
      case "category":
        console.log(event.target.value);
        setCategory(event.target.value);
        break;
      case "mean":
        console.log(event.target.value);
        setMean(event.target.value);
        break;
      case "any":
        console.log(event.target.value);
        setAny(event.target.value);
        break;
      default:
        alert("key not found");
    }
  };

  const handleSubmit = (event: React.ChangeEvent<HTMLInputElement>) => {
    event.preventDefault();
    addWord(word, category, mean, any);
  };

  return (
    <div className="App">
      <div>
        <Paper className={classes.paper}>
          <form className={classes.root} onSubmit={handleSubmit}>
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
                  label="単語"
                  variant="outlined"
                  onChange={handleChange}
                  value={word}
                />
              </Grid>

              <Grid item>
                <TextField
                  id="outlined-select-currency"
                  name="category"
                  select
                  label="種類"
                  value={category}
                  onChange={handleChange}
                  variant="outlined"
                >
                  {categories.map(option => (
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
                  label="意味"
                  variant="outlined"
                  onChange={handleChange}
                  value={mean}
                />
              </Grid>

              <Grid item>
                <TextField
                  id="outlined-multiline-static"
                  name="any"
                  label="補足"
                  multiline
                  rows="4"
                  variant="outlined"
                  onChange={handleChange}
                  value={any}
                />
              </Grid>

              <Grid item>
                <Button
                  type="submit"
                  size="medium"
                  variant="contained"
                  color="primary"
                >
                  単語を追加
                </Button>
              </Grid>
            </Grid>
          </form>
        </Paper>
      </div>
    </div>
  );
}

export default App;

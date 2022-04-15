import {useRef} from "react";

import Card from "../ui/Card";
import classes from "./NewMovieForm.module.css";

import {styled} from "@mui/material/styles";
import {gql, useMutation, useQuery} from "@apollo/client";
import {Stack} from "@mui/material";
import Button from "@mui/material/Button";


function NewMovieForm() {
    const NEW_MOVIE = gql`
        mutation CreateMovie ($title: String!, $description: String!, $genre: String!, $rank: Int!, $director_id: ID) {
            createMovie(movie: {title: $title , description: $description, genre: $genre , rank: $rank, director_id: $director_id}) {
                id
            }
        }
    `;
    const DIRECTOR_ID = gql`
        query DirectorIdByName($director: String!) {
            directorIdByName(name: $director)
        }
    `;

    const titleInputRef = useRef();
    const imageInputRef = useRef();
    const directorInputRef = useRef();
    const descriptionInputRef = useRef();
    const reviewInputRef = useRef();
    const rankInputRef = useRef();
    const worthInputRef = useRef();
    const genreInputRef = useRef();

    const Input = styled("input")({
        display: "none",
    });

    const id = useQuery(DIRECTOR_ID,
        {
            variables: {
                director: directorInputRef.current?.value || 'no director',
            }
        }).data

    const [addMovie, {data, loading, error}] = useMutation(NEW_MOVIE,
        {
            variables: {
                title: titleInputRef.current?.value || 'Unknown',
                image: imageInputRef.current?.value || 'No Picture',
                description: descriptionInputRef.current?.value || 'No Description',
                review: reviewInputRef.current?.value || 'Doesnt Have Any Reviews',
                rank: rankInputRef.current?.value || 'No Rank Was Given',
                worth: worthInputRef.current?.value || 'No Worth Was Given',
                genre: genreInputRef.current?.value || 'No Genre Given',
                director_id: id.value,
            }
        });
    return (
        <Card>
            <form className={classes.form}>
                <div className={classes.control}>
                    <label htmlFor="title">Movie Title</label>
                    <input type="text" datatype="String" required id="title" ref={titleInputRef}/>
                </div>

                <div className={classes.im}>
                    <label htmlFor="image">Movie Image</label>
                    <input type="url" id="image" ref={imageInputRef}/>
                </div>

                <Stack direction="row" alignItems="center" spacing={2} className={classes.but}>
                    <label htmlFor="contained-button-file">
                        <Input
                            accept="image/*"
                            multiple
                            type="file"
                            id="contained-button-file"
                            ref={imageInputRef}
                        />
                        <Button variant="contained" component="span">
                            Upload
                        </Button>
                    </label>
                </Stack>

                <div className={classes.control}>
                    <label htmlFor="director">Director's Name</label>
                    <input type="text" required id="director" ref={directorInputRef} datatype="String"/>
                </div>

                <div className={classes.control}>
                    <label htmlFor="worth">
                        How Much Do You Think This Movie Is worth Waching?
                    </label>
                    <input type="range" datatype="Int" id="worth" min="1" max="5" ref={worthInputRef}/>
                </div>

                <div className={classes.control}>
                    <label htmlFor="genre">What is the genre of this movie?</label>
                    <select name="genre" id="genre" ref={genreInputRef} required datatype="String">
                        <option value="action">Action</option>
                        <option value="drama">Drama</option>
                        <option value="comedy">Comedy</option>
                        <option value="crime">Crime</option>
                        <option value="animation">Animation</option>
                        <option value="fantasy">Fantasy</option>
                        <option value="romance">Romance</option>
                        <option value="thriller">Thriller</option>
                        <option value="horror">Horror</option>
                        <option value="science fiction">Science Fiction</option>
                        <option value="historical">Historical</option>
                        <option value="western">Western</option>
                    </select>
                </div>

                <div className={classes.control}>
                    <label htmlFor="description">Description</label>
                    <textarea
                        id="description"
                        type="text"
                        datatype="String"
                        required
                        rows="5"
                        ref={descriptionInputRef}
                    ></textarea>
                </div>

                <div className={classes.control}>
                    <label htmlFor="review">Review</label>
                    <textarea id="review" rows="5" datatype="String" ref={reviewInputRef}></textarea>
                </div>

                <div className={classes.ctrl}>
                    <label htmlFor="rank">Add Your Rank</label>
                    <input
                        type="number"
                        name="ranking"
                        id="ranking"
                        min="0"
                        max="100"
                        ref={rankInputRef}
                        datatype="Int"
                    ></input>
                </div>

                <div className={classes.actions}>
                    <button onClick={addMovie} type="button">Add Movie</button>
                </div>
            </form>
        </Card>
    );
}

export default NewMovieForm;
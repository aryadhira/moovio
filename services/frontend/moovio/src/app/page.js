import MovieList from "@/components/movielist";

const Home = async () => {
 
  const GetMovieList = async (type) => {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/movies/getmovielist?list=${type}`)
      const movies = await res.json()

      return movies
  }


  const latestdata = await GetMovieList("latest")
  const imdbdata = await GetMovieList("imdb")

  return (
    <div>
      <MovieList moviedata={latestdata} header={"Latest Uploaded Movies"}/>
      <MovieList moviedata={imdbdata} header={"Top Rating Movies"}/>
    </div>
  );
}

export default Home
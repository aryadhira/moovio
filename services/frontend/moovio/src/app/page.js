import MovieList from "@/components/movielist";

const Home = async () => {
 
  const GetMovieList = async () => {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/getmovielist`)
      const movies = await res.json()

      return movies
  }

  const data = await GetMovieList()

  return (
    <MovieList moviedata={data}/>
  );
}

export default Home
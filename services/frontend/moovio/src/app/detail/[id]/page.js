"use client"
import VideoPlayer from '@/components/videoplayer';
import { useState, useEffect } from 'react'

const Modal = ({ onClose, children }) => {
    const handleClose = () => {
        // Make a request to stop the stream using fetch
        fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/streamer/stopstream`, {
            method: 'POST',
        })
        .then(response => {
            if (response.ok) {
                console.log('Stream stopped successfully');
            } else {
                console.error('Error stopping stream:', response.status);
            }
        })
        .catch(error => {
            console.error('Error stopping stream:', error);
        });

        // Close the modal
        onClose();
    };

    return (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center">
            <div className="bg-white p-5 rounded-lg">
                {children}
                <button className="mt-4 bg-red-500 text-white px-4 py-2 rounded hover:bg-red-400" onClick={handleClose}>Close</button>
            </div>
        </div>
    );
}

const Page = ({ params }) => {
    const fallbackImage = 'https://placehold.co/600x900?text=Cover+Not+Available'

    const handleImgError = (e) => {
        e.target.src = fallbackImage; // Change the src to a fallback image
    };

    const [data,setData] = useState();
    const [loading, setLoading] = useState(true);
    const [showModal, setShowModal] = useState(false);
    const [modalContent, setModalContent] = useState('');
    const [parammovie, setParamMovie] = useState()

    useEffect(() => {
        fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/movies/getmoviedetail?id=${params.id}`).
        then((res)=> res.json()).
        then((data) => {
            setData(data.data)
            setLoading(false)
        })
    },[])

    if(loading) {
        return (
            <div className="flex justify-center items-center h-screen">Loading...</div>
        );
    }

    const formatGenres = (genres) => {
        return genres.join(" - ");
    };

    const openModal = (quality) => {
        setModalContent(`${quality}`);
        setParamMovie({
            title : `${data.title}`,
            quality : `${quality}`,
        });
        setShowModal(true);
    }


    return (
        <div className="w-full p-5 flex md:flex-row sm:flex-col">
            <div className="w-1/2 flex justify-center">
                <img width={600} height={600} src={data.cover} alt="..." onError={handleImgError}/>
            </div>
            <div className="p-10 w-1/2">
                <h1 className="font-bold text-2xl">{data.title}</h1>
                <h3 className="pt-5 text-red-500 text-md font-bold">Release Year : {data.year}</h3>
                <h3 className="pt-5 text-red-500 text-md font-bold">Genre : {formatGenres(data.category)}</h3>
                <h3 className="pt-5 pb-5 text-red-500">⭐️ {data.rating}</h3>
                <h2 className="font-bold text-xl">Synopsis:</h2>
                <p className="w-3/4 text-sm italic pt-5 pb-5">{data.synopsis}</p>
                <div className="flex gap-5">
                    <button className="rounded p-2 bg-red-500 w-20 shadow-xl hover:bg-red-400 hover:transition duration-300 ease-in-out hover:scale-125" onClick={() => openModal('720p')}>720p</button>
                    <button className="rounded p-2 bg-red-500 w-20 shadow-xl hover:bg-red-400 hover:transition duration-300 ease-in-out hover:scale-125" onClick={() => openModal('1080p')}>1080p</button>
                </div>
            </div>
            {showModal && (
                <Modal onClose={() => setShowModal(false)}>
                    <p className='text-slate-500'>{data.title} {modalContent}</p>
                    <VideoPlayer  param={parammovie}/>
                </Modal>
            )}
        </div>
        
    );
}

export default Page;
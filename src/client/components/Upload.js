import Image from "next/image";

const Upload = () => {
    return (
        <>
            <button className="hidden place-items-center px-2 border-2 rounded-md bg-gray-100 md:flex">
                <Image
                    src="/icons/plus.svg"
                    alt="Upload Icon"
                    width={24}
                    height={24}
                />
                <span className="md:m-2">Upload</span>
            </button>

            <button className="fixed bottom-20 right-6 border-2 p-2 rounded-md bg-gray-100 md:hidden">
                <Image
                    src="/icons/plus.svg"
                    alt="Plus Icon"
                    width={24}
                    height={24}
                ></Image>
            </button>
        </>
    );
};

export default Upload;

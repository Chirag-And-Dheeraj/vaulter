import Image from "next/image";

const Files = () => {
    return (
        <main>
            <button className="fixed bottom-20 right-8 border-2 p-2 rounded-md bg-gray-100" >
                <Image
                    src="/icons/plus.svg"
                    alt="Plus Icon"
                    width={24}
                    height={24}
                ></Image>
            </button>
        </main>
    );
};

export default Files;

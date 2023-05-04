import Image from "next/image";

const Item = ({ src, alt, name }) => {
    return (
        <div className="flex flex-col place-items-center">
            <Image
                src={src}
                alt={alt}
                width={24}
                height={24}
            ></Image>
            {name}
        </div>
    );
};

export default Item;

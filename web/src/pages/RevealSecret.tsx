import React, { useState } from "react";
import {
  Button,
  useToast,
} from "@chakra-ui/react";
import { Card, CardBody } from "@chakra-ui/react";
import { CopyIcon, UnlockIcon } from "@chakra-ui/icons";

export const RevealSecret = () => {
  const [isRevealed, setIsRevealed] = useState(false);
  const toast = useToast();

  return (
    <Card variant="elevated">
      <CardBody>
        {isRevealed === false && (
          <>
            <p className="text-xl my-8">
              Be aware! The following secret can only be revealed one time.
            </p>
            <Button
              leftIcon={<UnlockIcon />}
              colorScheme="orange"
              onClick={() => {
                setIsRevealed(true);
                toast({
                  title: "Secret revealed!",
                  status: "success",
                  duration: 4000,
                  isClosable: true,
                });
              }}
            >
              Reveal secret
            </Button>
          </>
        )}
        {isRevealed === true && (
          <>
            <p id="secret-content" className="text-xl my-8">
              asd asdadsa das
            </p>
            <Button
              leftIcon={<CopyIcon />}
              colorScheme="orange"
              onClick={() => {
                const secretContent = document.getElementById("secret-content");
                navigator.clipboard.writeText(secretContent?.textContent || "");
                toast({
                  title: "Secret copied!",
                  status: "success",
                  duration: 6000,
                  isClosable: true,
                });
              }}
            >
              Copy secret
            </Button>
          </>
        )}
      </CardBody>
    </Card>
  );
};

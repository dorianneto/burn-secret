import React, { useEffect, useState } from "react";
import { Button, useToast } from "@chakra-ui/react";
import { Card, CardBody } from "@chakra-ui/react";
import { CopyIcon, UnlockIcon } from "@chakra-ui/icons";
import axios from "axios";
import { useNavigate, useParams } from "react-router-dom";

export const RevealSecret = () => {
  const [secret, setSecret] = useState(null);
  const [isRevealed, setIsRevealed] = useState(false);

  const { id } = useParams();
  const navigate = useNavigate()
  const toast = useToast();

  useEffect(() => {
    const fetchData = async () => {
      const { data, status } = await axios.get(
        `http://localhost/api/v1/secret/${id}`
      );

      if (status != 200) {
        toast({
          title: "Something goes wrong",
          status: "error",
          duration: 6000,
          isClosable: true,
        });
        return;
      }

      setSecret(data.data);
    };

    fetchData().catch(() => {
      navigate("/", { replace: true })
    });
  }, []);

  useEffect(() => {
    if (!isRevealed) {
      return;
    }

    const fetchData = async () => {
      await axios.delete(`http://localhost/api/v1/secret/${id}/burn`);
    };

    fetchData().catch(console.error);
  }, [isRevealed]);

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
              {secret}
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

import React from "react";
import styled from "styled-components";
import useSWR from "swr";

const fetcher = async () => await fetch("/api/tasks").then((res) => res.json());
export default function HelloPage() {
  const { data, isLoading, error } = useSWR("task", fetcher);

  if (isLoading) return <HelloUI>로딩 중.../</HelloUI>;
  if (error) return <HelloUI>에러...!</HelloUI>;

  return <HelloUI>데이터: {String(JSON.stringify(data))}</HelloUI>;
}

const HelloUI = styled.div`
  color: black;
`;

package com.url.shortener;

import org.springframework.data.repository.CrudRepository;
import org.springframework.data.jpa.repository.Query; 

import com.url.shortener.Url;

// This will be AUTO IMPLEMENTED by Spring into a Bean called urlRepository
// CRUD refers Create, Read, Update, Delete

public interface UrlRepository extends CrudRepository<Url, Integer> {

    Url findByUrl(String url);

    // ShortId -> find Url model
    Url findByShortId(String shortId);
}
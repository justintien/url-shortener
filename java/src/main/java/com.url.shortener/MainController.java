package com.url.shortener;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.validation.BindingResult;
import org.springframework.validation.ObjectError;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.URL;
import java.net.MalformedURLException;
import javax.servlet.http.HttpServletResponse;
import java.util.Date;

@Controller
// @RequestMapping(path="/")
public class MainController {
	@Autowired

	private UrlRepository urlRepository;
    private static final Logger logger = LoggerFactory.getLogger(MainController.class);

	@GetMapping(path="/")
	public @ResponseBody String index () {
		return "server alive!";
	}

	@PostMapping(path="/shorten")
	public @ResponseBody Url shorten (@RequestBody Url model, BindingResult bindingResult) {
       if (!isUrlValid(model.getUrl())) {
        	bindingResult.addError(new ObjectError("url", "Invalid url format: " + model.getUrl()));
			return new Url();
        }

		Url find = urlRepository.findByUrl(model.getUrl());
		logger.debug("find:{}", find);
		if (find != null)
			return find;
		Date date = new Date();
		String time36 = Long.toString(date.getTime(), 36).toLowerCase();
		String suffix4 = "0000" + Long.toString((int) (Math.random() * Math.pow(36, 4)), 36);
		suffix4 = suffix4.substring(suffix4.length() - 4, suffix4.length());
		String shortId = time36 + suffix4;

		logger.debug("time:{}, time36:{}, suffix4:{}", date.getTime(), time36, suffix4);
		
		model.setShortId(shortId);
		urlRepository.save(model);
		return model;
	}

	@GetMapping(path="/{shortId}")
	public @ResponseBody void redirect (@PathVariable("shortId") String shortId, HttpServletResponse resp) throws Exception {
		Url find = urlRepository.findByShortId(shortId);

		if (find == null) {
			resp.sendError(HttpServletResponse.SC_NOT_FOUND);
			return;
		}
		resp.addHeader("Location", find.getUrl());
		resp.setStatus(HttpServletResponse.SC_MOVED_PERMANENTLY);
	}

	@GetMapping(path="/list")
	public @ResponseBody Iterable<Url> getList() {
		return urlRepository.findAll();
	}

    private boolean isUrlValid(String url) {
        boolean valid = true;
        try {
            new URL(url);
        } catch (MalformedURLException e) {
			logger.error("isUrlValid MalformedURLException", e);

            valid = false;
        }
        return valid;
    }
}
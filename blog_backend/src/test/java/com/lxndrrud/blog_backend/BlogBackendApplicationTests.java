package com.lxndrrud.blog_backend;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;

@WebMvcTest
class BlogBackendApplicationTests {
	@Autowired
	private MockMvc mockMvc;

	@Test
	void testGetAllPosts() throws Exception {
		mockMvc.perform(MockMvcRequestBuilders.get("/posts/"))
				.andExpect(MockMvcResultMatchers.status().isOk())
				.andExpect(MockMvcResultMatchers.content().json("[" +
						"\"kek\"" +
						"]"));
	}

}
